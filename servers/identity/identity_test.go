// © Copyright 2024 Hewlett Packard Enterprise Development LP

// Package identity
package identity

import (
	"context"
	stdlog "log"
	"os"
	"reflect"
	"testing"

	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"
	cosi "sigs.k8s.io/container-object-storage-interface/proto"
)

func TestNew(t *testing.T) {
	type args struct {
		name   string
		logger logr.Logger
	}
	log := stdr.New(stdlog.New(os.Stdout, "", stdlog.LstdFlags))

	tests := []struct {
		name string
		args args
		want *Server
	}{
		{
			name: "Empty provisioner name",
			args: args{
				name:   "",
				logger: log,
			},
			want: nil,
		},
		{
			name: "Provisioner name present",
			args: args{
				name:   "test.cosi.driver",
				logger: log,
			},
			want: &Server{
				name: "test.cosi.driver",
				log:  log,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil && test.want == nil {
					t.Errorf("Expected a panic")
				}
			}()
			if got := New(test.args.name, test.args.logger); !reflect.DeepEqual(got, test.want) {
				t.Errorf("New() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestServer_DriverGetInfo(t *testing.T) {
	type fields struct {
		log  logr.Logger
		name string
	}
	type args struct {
		ctx context.Context
		req *cosi.DriverGetInfoRequest
	}
	log := stdr.New(stdlog.New(os.Stdout, "", stdlog.LstdFlags))

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *cosi.DriverGetInfoResponse
		wantErr bool
	}{
		{
			name: "Empty provisioner name",
			fields: fields{
				log:  log,
				name: "",
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverGetInfoRequest{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Provisioner name present",
			fields: fields{
				log:  log,
				name: "test.cosi.driver",
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverGetInfoRequest{},
			},
			want: &cosi.DriverGetInfoResponse{
				Name: "test.cosi.driver",
			},
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := &Server{
				log:  test.fields.log,
				name: test.fields.name,
			}
			got, err := s.DriverGetInfo(test.args.ctx, test.args.req)
			if (err != nil) != test.wantErr {
				t.Errorf("Server.DriverGetInfo() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Server.DriverGetInfo() = %v, want %v", got, test.want)
			}
		})
	}
}
