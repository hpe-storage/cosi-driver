// © Copyright 2024 Hewlett Packard Enterprise Development LP

// Package identity
package identity

import (
	"context"

	"github.com/go-logr/logr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	cosi "sigs.k8s.io/container-object-storage-interface/proto"
)

// Server implements cosi.IdentityServer interface.
type Server struct {
	log  logr.Logger
	name string
	cosi.UnsafeIdentityServer
}

// Interface guards.
var _ cosi.IdentityServer = (*Server)(nil)

// New returns identity.Server with name set to the "name" argument.
func New(name string, logger logr.Logger) *Server {
	if name == "" {
		panic("Object storage provisioner name has not been provided.")
	}

	return &Server{
		log:  logger,
		name: name,
	}
}

// DriverGetInfo call is meant to retrieve the unique provisioner Identity.
func (s *Server) DriverGetInfo(_ context.Context, _ *cosi.DriverGetInfoRequest) (*cosi.DriverGetInfoResponse, error) {
	if s.name == "" {
		return nil, status.Error(codes.NotFound, "DriverName is empty")
	}

	s.log.Info("COSI driver info. ", "name", s.name)

	return &cosi.DriverGetInfoResponse{
		Name: s.name,
	}, status.Error(codes.OK, "OK")
}
