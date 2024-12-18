// © Copyright 2024 Hewlett Packard Enterprise Development LP

// Package utils
package utils

import (
	"context"
	"errors"
	stdlog "log"
	"os"
	"reflect"
	"testing"

	gomonkey "github.com/agiledragon/gomonkey/v2"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"
)

func TestNewS3Client(t *testing.T) {
	type args struct {
		ctx       context.Context
		accessKey string
		secretKey string
		endpoint  string
	}
	ss := &session.Session{}
	output := []gomonkey.OutputCell{
		{Values: gomonkey.Params{ss, nil}},
		{Values: gomonkey.Params{nil, errors.New("failed to create new session")}},
	}
	patches := gomonkey.ApplyFuncSeq(session.NewSession, output)
	defer patches.Reset()
	s := &s3.S3{}
	patches.ApplyFuncReturn(s3.New, s)
	tests := []struct {
		name    string
		args    args
		want    *S3Client
		wantErr bool
	}{
		{
			name: "Create s3 client successfully",
			args: args{
				ctx:       context.TODO(),
				accessKey: "testuser",
				secretKey: "testuser",
				endpoint:  "http://10.157.37.124:8080",
			},
			want: &S3Client{
				Client: s,
			},
			wantErr: false,
		},
		{
			name: "Fails to new session",
			args: args{
				ctx:       context.TODO(),
				accessKey: "testuser",
				secretKey: "testuser",
				endpoint:  "http://10.157.37.124:8080",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewS3Client(tt.args.ctx, tt.args.accessKey, tt.args.secretKey, tt.args.endpoint)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewS3Client() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewS3Client() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestS3Client_CreateBucket(t *testing.T) {
	type fields struct {
		Client s3iface.S3API
	}
	type args struct {
		ctx        context.Context
		bucketName string
		param      map[string]string
	}

	s3c := &s3.S3{}
	awsErr1 := awserr.New(s3.ErrCodeBucketAlreadyExists, "", errors.New("BucketAlreadyExists"))
	awsErr2 := awserr.New(s3.ErrCodeBucketAlreadyOwnedByYou, "", errors.New("BucketAlreadyOwnedByYou"))
	awsErr3 := awserr.New("RequestError", "", errors.New("connect: connection refused"))

	outputSeqCreate := []gomonkey.OutputCell{
		{Values: gomonkey.Params{nil, nil}},
		{Values: gomonkey.Params{nil, error(awsErr1)}},
		{Values: gomonkey.Params{nil, error(awsErr2)}},
		{Values: gomonkey.Params{nil, error(awsErr3)}},
		{Values: gomonkey.Params{nil, errors.New("failed to create bucket")}},
		{Values: gomonkey.Params{nil, nil}},
		{Values: gomonkey.Params{nil, nil}},
		{Values: gomonkey.Params{nil, nil}},
		{Values: gomonkey.Params{nil, nil}},
		{Values: gomonkey.Params{nil, nil}},
		{Values: gomonkey.Params{nil, nil}},
		{Values: gomonkey.Params{nil, nil}},
	}
	patches := gomonkey.ApplyMethodSeq(s3c, "CreateBucket", outputSeqCreate)
	defer patches.Reset()

	outputSeqPutTag := []gomonkey.OutputCell{
		{Values: gomonkey.Params{nil, nil}},
		{Values: gomonkey.Params{nil, errors.New("failed to put bucket tag config")}},
		{Values: gomonkey.Params{nil, nil}},
		{Values: gomonkey.Params{nil, nil}},
	}
	patches.ApplyMethodSeq(s3c, "PutBucketTagging", outputSeqPutTag)

	outputSeqDel := []gomonkey.OutputCell{
		{Values: gomonkey.Params{nil, errors.New("failed to delete bucket")}},
		{Values: gomonkey.Params{nil, nil}},
		{Values: gomonkey.Params{nil, errors.New("failed to delete bucket")}},
	}
	patches.ApplyMethodSeq(s3c, "DeleteBucket", outputSeqDel)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Create bucket successfully",
			fields: fields{
				Client: s3c,
			},
			args: args{
				ctx:        context.TODO(),
				bucketName: "bucket1",
				param:      make(map[string]string),
			},
			wantErr: false,
		},
		{
			name: "Bucket already exists",
			fields: fields{
				Client: s3c,
			},
			args: args{
				ctx:        context.TODO(),
				bucketName: "bucket1",
				param:      make(map[string]string),
			},
			wantErr: false,
		},
		{
			name: "Bucket already owned",
			fields: fields{
				Client: s3c,
			},
			args: args{
				ctx:        context.TODO(),
				bucketName: "bucket1",
				param:      make(map[string]string),
			},
			wantErr: false,
		},
		{
			name: "Failure due to connection error",
			fields: fields{
				Client: s3c,
			},
			args: args{
				ctx:        context.TODO(),
				bucketName: "bucket1",
				param:      make(map[string]string),
			},
			wantErr: true,
		},
		{
			name: "Failed to create bucket",
			fields: fields{
				Client: s3c,
			},
			args: args{
				ctx:        context.TODO(),
				bucketName: "bucket1",
				param:      make(map[string]string),
			},
			wantErr: true,
		},
		{
			name: "Apply tags to the bucket successfully",
			fields: fields{
				Client: s3c,
			},
			args: args{
				ctx:        context.TODO(),
				bucketName: "bucket1",
				param:      map[string]string{bucketTagsKey: "key1=val1"},
			},
			wantErr: false,
		},
		{
			name: "Failed to put bucket tag(s)",
			fields: fields{
				Client: s3c,
			},
			args: args{
				ctx:        context.TODO(),
				bucketName: "bucket1",
				param:      map[string]string{bucketTagsKey: "key1=val1"},
			},
			wantErr: true,
		},
		{
			name: "Empty bucket tag parameter value",
			fields: fields{
				Client: s3c,
			},
			args: args{
				ctx:        context.TODO(),
				bucketName: "bucket1",
				param:      map[string]string{bucketTagsKey: ""},
			},
			wantErr: false,
		},
		{
			name: "Invalid bucket key-value format",
			fields: fields{
				Client: s3c,
			},
			args: args{
				ctx:        context.TODO(),
				bucketName: "bucket1",
				param:      map[string]string{bucketTagsKey: "keyvalue"},
			},
			wantErr: true,
		},
		{
			name: "Invalid bucket tag key - bucket delete fails",
			fields: fields{
				Client: s3c,
			},
			args: args{
				ctx:        context.TODO(),
				bucketName: "bucket1",
				param:      map[string]string{bucketTagsKey: "=val"},
			},
			wantErr: true,
		},
		{
			name: "Empty strings in the list of bucket tags",
			fields: fields{
				Client: s3c,
			},
			args: args{
				ctx:        context.TODO(),
				bucketName: "bucket1",
				param:      map[string]string{bucketTagsKey: "key=val,k=v,,"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &S3Client{
				Client: tt.fields.Client,
			}
			if err := s.CreateBucket(tt.args.ctx, tt.args.bucketName, tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("S3Client.CreateBucket() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestS3Client_DeleteBucket(t *testing.T) {
	type fields struct {
		Client s3iface.S3API
	}
	type args struct {
		ctx        context.Context
		bucketName string
	}

	s3c := &s3.S3{}

	outputSeq := []gomonkey.OutputCell{
		{Values: gomonkey.Params{nil, nil}},
		{Values: gomonkey.Params{nil, errors.New("failed to delete bucket")}},
	}
	patch := gomonkey.ApplyMethodSeq(s3c, "DeleteBucket", outputSeq)
	defer patch.Reset()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Deleted bucket successfully",
			fields: fields{
				Client: s3c,
			},
			args: args{
				ctx:        context.TODO(),
				bucketName: "bucket1",
			},
			wantErr: false,
		},
		{
			name: "Failed to delete bucket",
			fields: fields{
				Client: s3c,
			},
			args: args{
				ctx:        context.TODO(),
				bucketName: "bucket1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &S3Client{
				Client: tt.fields.Client,
			}
			if err := s.DeleteBucket(tt.args.ctx, tt.args.bucketName); (err != nil) != tt.wantErr {
				t.Errorf("S3Client.DeleteBucket() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetLoggerFromContext(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	log := stdr.New(stdlog.New(os.Stdout, "", stdlog.LstdFlags))
	tests := []struct {
		name   string
		args   args
		ctxVal *logr.Logger
	}{
		{
			name: "Context contains loggerKey value",
			args: args{
				ctx: context.WithValue(context.TODO(), LoggerKey, &log),
			},
		},
		{
			name: "Context without loggerKey value",
			args: args{
				ctx: context.Background(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLoggerFromContext(tt.args.ctx); got == nil {
				t.Errorf("Expected non-nil pointer to logger")
			}
		})
	}
}
