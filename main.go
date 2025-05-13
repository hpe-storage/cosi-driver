// © Copyright 2024 Hewlett Packard Enterprise Development LP

package main

import (
	"context"
	"errors"
	"fmt"
	"hpe-cosi-osp/servers/identity"
	"hpe-cosi-osp/servers/provisioner"
	"hpe-cosi-osp/servers/provisioner/utils"
	"io/fs"
	"net/url"

	stdlog "log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
	cosi "sigs.k8s.io/container-object-storage-interface/sidecar/pkg/provisioner"
)

const (
	// Name of the COSI driver
	ospName = "cosi.hpe.com"
	// Address of the socket connection between the sidecar and COSI driver
	driverAddress = "unix:///var/lib/cosi/cosi.sock"
)

var (
	log logr.Logger
)

// Verbosity levels of the logr logger
const (
	debugVerbosity = 4
	infoVerbosity  = 2
	warnVerbosity  = 1
	errorVerbosity = 0
)

// interceptorLogger adapts logr logger to interceptor logger for the gRPC server
func interceptorLogger(l logr.Logger) logging.Logger {
	return logging.LoggerFunc(func(_ context.Context, lvl logging.Level, msg string, fields ...any) {
		l := l.WithValues(fields...)
		switch lvl {
		case logging.LevelDebug:
			l.V(debugVerbosity).Info(msg)
		case logging.LevelInfo:
			l.V(infoVerbosity).Info(msg)
		case logging.LevelWarn:
			l.V(warnVerbosity).Info(msg)
		case logging.LevelError:
			l.V(errorVerbosity).Info(msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}

func main() {
	// Initialize logger
	stdr.SetVerbosity(2)
	// The stdlog flags enable date, time (including microseconds) and filename
	log = stdr.New(stdlog.New(os.Stdout, "", stdlog.LstdFlags|stdlog.Lmicroseconds|stdlog.Lshortfile))

	// Create context and attach logger
	ctx := context.Background()
	ctx = context.WithValue(ctx, utils.LoggerKey, &log)

	// Get notified when an interrupt or termination signal is received
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// Create the identity and provisioner servers for the driver
	identityServer := identity.New(ospName, log)
	provisionerServer, err := provisioner.New(log)
	if err != nil {
		log.Error(err, "failed to create provisioner server")
		os.Exit(1)
	}

	// Define options to be used by the gRPC servers, to log the start and finish calls of a request
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	}
	grpcOptions := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(interceptorLogger(log), opts...),
		),
	}

	// Initialize the COSI servers
	server, err := cosi.NewCOSIProvisionerServer(driverAddress,
		identityServer,
		provisionerServer,
		grpcOptions)
	if err != nil {
		log.Error(err, "failed to initialize COSI servers")
		os.Exit(1)
	}

	// Remove the UNIX socket file if it already exists,
	// to avoid errors when the driver restarts after a crash
	url, err := url.Parse(driverAddress)
	if err != nil {
		log.Error(err, "failed to parse the UNIX socket address")
		os.Exit(1)
	}
	if url.Path != "" {
		if _, err := os.Stat(url.Path); !errors.Is(err, fs.ErrNotExist) {
			if err := os.RemoveAll(url.Path); err != nil {
				log.Error(err, "failed to remove socket")
				os.Exit(1)
			}
		}
	}

	// Register the COSI servers, start listening and serving requests on the socket
	err = server.Run(ctx)
	if err != nil {
		log.Error(err, "failed to run COSI servers")
		os.Exit(1)
	}
}
