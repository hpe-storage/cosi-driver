// © Copyright 2024 Hewlett Packard Enterprise Development LP

// Package provisioner
package provisioner

import (
	"context"
	"errors"
	"hpe-cosi-osp/servers/provisioner/test_utils"
	"hpe-cosi-osp/servers/provisioner/utils"
	stdlog "log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"

	"hpe-cosi-osp/iam"

	gomonkey "github.com/agiledragon/gomonkey/v2"
	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	v1alpha1 "sigs.k8s.io/container-object-storage-interface/client/apis/objectstorage/v1alpha1"
	bucketclientset "sigs.k8s.io/container-object-storage-interface/client/clientset/versioned"
	fakebucketclientset "sigs.k8s.io/container-object-storage-interface/client/clientset/versioned/fake"
	cosi "sigs.k8s.io/container-object-storage-interface/proto"
)

var secretName = "cosi-secret"
var namespace = "cosi-secret-ns"

// TestNew verifies the initialization of a new Server instance.
func TestNew(t *testing.T) {
	type args struct {
		logger logr.Logger
	}
	log := stdr.New(stdlog.New(os.Stdout, "", stdlog.LstdFlags))

	kubeConfig := &rest.Config{}
	cc := []gomonkey.OutputCell{
		{Values: gomonkey.Params{kubeConfig, nil}},
		{Values: gomonkey.Params{kubeConfig, errors.New("failed to get cluster config")}},
		{Values: gomonkey.Params{kubeConfig, nil}},
		{Values: gomonkey.Params{kubeConfig, nil}},
	}
	patches := gomonkey.ApplyFuncSeq(rest.InClusterConfig, cc)
	defer patches.Reset()
	k8sClientset := &kubernetes.Clientset{}
	kcs := []gomonkey.OutputCell{
		{Values: gomonkey.Params{k8sClientset, nil}},
		{Values: gomonkey.Params{k8sClientset, errors.New("failed to create k8s client set")}},
		{Values: gomonkey.Params{k8sClientset, nil}},
	}

	patches.ApplyFuncSeq(kubernetes.NewForConfig, kcs)
	bucketClientset := &bucketclientset.Clientset{}
	bcs := []gomonkey.OutputCell{
		{Values: gomonkey.Params{bucketClientset, nil}},
		{Values: gomonkey.Params{bucketClientset, errors.New("failed to create bucket client set")}},
	}
	patches.ApplyFuncSeq(bucketclientset.NewForConfig, bcs)

	tests := []struct {
		name    string
		args    args
		want    *Server
		wantErr bool
	}{
		{
			name: "Create new provisioner server successfully",
			args: args{
				logger: log,
			},
			want: &Server{
				log:             log,
				K8sClientset:    k8sClientset,
				BucketClientset: bucketClientset,
			},
			wantErr: false,
		},
		{
			name: "Fails to get cluster config",
			args: args{
				logger: log,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Fails to create k8s client set",
			args: args{
				logger: log,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Fails to create bucket client set",
			args: args{
				logger: log,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.logger)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestDriverCreateBucket tests the DriverCreateBucket method for creating a bucket.
// It ensures idempotency and validates the bucket creation logic.
func TestDriverCreateBucket(t *testing.T) {
	type fields struct {
		log             logr.Logger
		K8sClientset    *kubernetes.Clientset
		BucketClientset bucketclientset.Interface
	}
	type args struct {
		ctx context.Context
		req *cosi.DriverCreateBucketRequest
	}

	// Create patches to retrieve secret
	secret := createSecret(secretName, namespace, []byte("testuser"), []byte("testuser"), []byte("http://1.1.1.1:90"), nil)
	emptySecret := createSecret(secretName, namespace, []byte(""), []byte("testuser"), []byte("http://1.1.1.1:90"), nil)
	coreV1 := &corev1.CoreV1Client{}
	client := &kubernetes.Clientset{}
	mockSecret := test_utils.MockSecret{Secret: secret}
	var secretInt corev1.SecretInterface = mockSecret
	patches := gomonkey.ApplyMethodReturn(client, "CoreV1", coreV1)
	defer patches.Reset()
	patches.ApplyMethodReturn(coreV1, "Secrets", secretInt)
	mockSecrets := []gomonkey.OutputCell{
		{Values: gomonkey.Params{secret, nil}},
		{Values: gomonkey.Params{nil, errors.New("failed to get secret")}},
		{Values: gomonkey.Params{emptySecret, nil}},
		{Values: gomonkey.Params{secret, nil}},
		{Values: gomonkey.Params{secret, nil}},
	}
	patches.ApplyMethodSeq(secretInt, "Get", mockSecrets)

	// Patch createS3Client to return a mock S3Client
	mockS3c := &S3Client{
		BaseURL:    "http://mock-s3",
		HTTPClient: &http.Client{},
		AccessKey:  "testuser",
		SecretKey:  "testuser",
	}
	patches.ApplyFunc(createS3Client, func(ctx context.Context, parameters map[string]string, kcs *kubernetes.Clientset) (*S3Client, error) {
		return mockS3c, nil
	})
	// Patch CreateBucket to simulate success/failure for each test case
	patches.ApplyMethodSeq(mockS3c, "CreateBucket", []gomonkey.OutputCell{
		{Values: gomonkey.Params{nil}},                                   // Create bucket successfully
		{Values: gomonkey.Params{errors.New("missing parameter")}},       // Failure due to missing parameter
		{Values: gomonkey.Params{errors.New("failed to get secret")}},    // Failed to get secret
		{Values: gomonkey.Params{errors.New("invalid secret")}},          // Failure due to invalid secret
		{Values: gomonkey.Params{errors.New("failed to create bucket")}}, // Failed to create bucket
	})

	log := stdr.New(stdlog.New(os.Stdout, "", stdlog.LstdFlags))

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *cosi.DriverCreateBucketResponse
		wantErr bool
	}{
		{
			name: "Create bucket successfully",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: &bucketclientset.Clientset{},
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverCreateBucketRequest{
					Name: "bucket1",
					Parameters: map[string]string{
						"cosiUserSecretNamespace": namespace,
						"cosiUserSecretName":      secretName,
					},
				},
			},
			want: &cosi.DriverCreateBucketResponse{
				BucketId: "bucket1",
			},
			wantErr: false,
		},
		{
			name: "Failure due to missing parameter",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: &bucketclientset.Clientset{},
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverCreateBucketRequest{
					Name: "bucket1",
					Parameters: map[string]string{
						"cosiUserSecretName": secretName,
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed to get secret",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: &bucketclientset.Clientset{},
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverCreateBucketRequest{
					Name: "bucket1",
					Parameters: map[string]string{
						"cosiUserSecretNamespace": namespace,
						"cosiUserSecretName":      secretName,
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failure due to invalid secret",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: &bucketclientset.Clientset{},
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverCreateBucketRequest{
					Name: "bucket1",
					Parameters: map[string]string{
						"cosiUserSecretNamespace": namespace,
						"cosiUserSecretName":      secretName,
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed to create bucket",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: &bucketclientset.Clientset{},
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverCreateBucketRequest{
					Name: "bucket1",
					Parameters: map[string]string{
						"cosiUserSecretNamespace": namespace,
						"cosiUserSecretName":      secretName,
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				log:             tt.fields.log,
				K8sClientset:    tt.fields.K8sClientset,
				BucketClientset: tt.fields.BucketClientset,
			}
			got, err := s.DriverCreateBucket(tt.args.ctx, tt.args.req)
			s.log.Error(err, "test error")
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.DriverCreateBucket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.DriverCreateBucket() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestDriverDeleteBucket tests the DriverDeleteBucket method for deleting a bucket.
// It ensures idempotency and validates the bucket deletion logic.
func TestDriverDeleteBucket(t *testing.T) {
	type fields struct {
		log             logr.Logger
		K8sClientset    *kubernetes.Clientset
		BucketClientset bucketclientset.Interface
	}
	type args struct {
		ctx context.Context
		req *cosi.DriverDeleteBucketRequest
	}
	// Create logger
	log := stdr.New(stdlog.New(os.Stdout, "", stdlog.LstdFlags))

	// Mock secret and bucket methods
	bucket := &v1alpha1.Bucket{
		ObjectMeta: metav1.ObjectMeta{
			Name: "bucket1",
		},
		Spec: v1alpha1.BucketSpec{
			Parameters: map[string]string{
				"cosiUserSecretNamespace": namespace,
				"cosiUserSecretName":      secretName,
			},
		},
	}
	bcs := fakebucketclientset.NewSimpleClientset(bucket)

	// Create mock secret instances
	secret := createSecret(secretName, namespace, []byte("testuser"), []byte("testuser"), []byte("http://1.1.1.1:90"), nil)
	emptySecret := createSecret(secretName, namespace, []byte(""), []byte("testuser"), []byte("http://1.1.1.1:90"), nil)
	coreV1 := &corev1.CoreV1Client{}
	client := &kubernetes.Clientset{}
	mockSecret := test_utils.MockSecret{Secret: secret}
	var secretInt corev1.SecretInterface = mockSecret

	// Create patches to retrieve secret
	patches := gomonkey.ApplyMethodReturn(client, "CoreV1", coreV1)
	defer patches.Reset()
	patches.ApplyMethodReturn(coreV1, "Secrets", secretInt)
	mockSecrets := []gomonkey.OutputCell{
		{Values: gomonkey.Params{secret, nil}},
		{Values: gomonkey.Params{nil, errors.New("failed to get secret")}},
		{Values: gomonkey.Params{emptySecret, nil}},
		{Values: gomonkey.Params{secret, nil}},
		{Values: gomonkey.Params{secret, nil}},
	}
	patches.ApplyMethodSeq(secretInt, "Get", mockSecrets)

	// Patch createS3Client to return a mock S3Client
	mockS3c := &S3Client{
		BaseURL:    "http://mock-s3",
		HTTPClient: &http.Client{},
		AccessKey:  "testuser",
		SecretKey:  "testuser",
	}
	patches.ApplyFunc(createS3Client, func(ctx context.Context, parameters map[string]string, kcs *kubernetes.Clientset) (*S3Client, error) {
		return mockS3c, nil
	})
	// Patch DeleteBucket to simulate success/failure for each test case
	patches.ApplyMethodSeq(mockS3c, "DeleteBucket", []gomonkey.OutputCell{
		{Values: gomonkey.Params{nil}},                                   // Deleted bucket successfully
		{Values: gomonkey.Params{errors.New("failed to get secret")}},    // Failed to get secret
		{Values: gomonkey.Params{errors.New("invalid secret")}},          // Failure due to invalid secret
		{Values: gomonkey.Params{errors.New("failed to delete bucket")}}, // Failed to delete bucket
	})

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *cosi.DriverDeleteBucketResponse
		wantErr bool
	}{
		{
			name: "Deleted bucket successfully",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: bcs,
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverDeleteBucketRequest{
					BucketId: "bucket1",
				},
			},
			want:    &cosi.DriverDeleteBucketResponse{},
			wantErr: false,
		},
		{
			name: "Failed to get secret",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: bcs,
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverDeleteBucketRequest{
					BucketId: "bucket1",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failure due to invalid secret",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: bcs,
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverDeleteBucketRequest{
					BucketId: "bucket1",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed to delete bucket",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: bcs,
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverDeleteBucketRequest{
					BucketId: "bucket1",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				log:             tt.fields.log,
				K8sClientset:    tt.fields.K8sClientset,
				BucketClientset: tt.fields.BucketClientset,
			}
			got, err := s.DriverDeleteBucket(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.DriverDeleteBucket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.DriverDeleteBucket() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestDriverGrantBucketAccess tests the DriverGrantBucketAccess method for granting access to a bucket.
// It validates the creation of credentials and access policies.
func TestDriverGrantBucketAccess(t *testing.T) {
	type fields struct {
		log             logr.Logger
		K8sClientset    *kubernetes.Clientset
		BucketClientset bucketclientset.Interface
	}
	type args struct {
		ctx context.Context
		req *cosi.DriverGrantBucketAccessRequest
	}

	// Create patches to retrieve secret
	systemId := "DUMMY_SERIAL_NUMBER"
	glcpCreds := getIAMCredentials()
	secret := createSecret(secretName, namespace, []byte("testuser"), []byte("testuser"), []byte("http://1.1.1.1:90"), &glcpCreds)
	coreV1 := &corev1.CoreV1Client{}
	client := &kubernetes.Clientset{}
	mockSecret := test_utils.MockSecret{Secret: secret}
	var secretInt corev1.SecretInterface = mockSecret
	patches := gomonkey.ApplyMethodReturn(client, "CoreV1", coreV1)
	defer patches.Reset()
	patches.ApplyMethodReturn(coreV1, "Secrets", secretInt)
	log := stdr.New(stdlog.New(os.Stdout, "", stdlog.LstdFlags))
	accessName := "bucket1_user"
	bucketName := "bucket1"
	policy := iam.NewS3Policy(utils.ACCESS_POLICY_PREFIX+accessName, bucketName, systemId, nil, context.Background())
	user := iam.NewS3User(utils.USER_PREFIX+accessName, systemId, nil, context.Background())

	removeSecretData := func(remField string) *v1.Secret {
		data := map[string][]byte{utils.GLCP_USER_CLIENTID: []byte(glcpCreds.GLCPUser),
			utils.GLCP_USER_SECRET_KEY: []byte(glcpCreds.GLCPUserSecretKey),
			utils.DSCC_ZONE:            []byte(glcpCreds.DSCCZone),
			utils.ALLETRA_MP_X10K_SNO:  []byte(glcpCreds.SystemId),
		}
		delete(data, remField)
		return &v1.Secret{Data: data}
	}

	tests := []struct {
		name         string
		fields       fields
		args         args
		want         *cosi.DriverGrantBucketAccessResponse
		setupPatches func() *gomonkey.Patches
		wantErr      bool
	}{
		{
			name: "Create bucket access successfully",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: &bucketclientset.Clientset{},
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverGrantBucketAccessRequest{
					Name:               accessName,
					AuthenticationType: cosi.AuthenticationType_Key,
					Parameters: map[string]string{
						"cosiUserSecretNamespace": namespace,
						"cosiUserSecretName":      secretName,
					},
				},
			},
			want: &cosi.DriverGrantBucketAccessResponse{
				AccountId: accessName,
				Credentials: map[string]*cosi.CredentialDetails{
					"s3": {Secrets: map[string]string{
						"accessKeyID":     "user_bucket1_user",
						"accessSecretKey": "user_bucket1_user_key",
						"endpoint":        "http://1.1.1.1:90",
						"region":          "us-east-1",
					}},
				},
			},
			setupPatches: func() *gomonkey.Patches {
				p := gomonkey.ApplyFuncReturn(getSecret, secret, nil)
				p = p.ApplyFuncReturn(os.Getenv, "DUMMY_PROXY_URL")
				p = p.ApplyFuncReturn(getAccessToken, "bearerdummyoxyzxxzzz12xxxx341111zzzzyyyyyyQQQQQHHHHH", nil)
				p = p.ApplyFuncReturn(iam.NewS3Policy, policy)
				p = p.ApplyMethodReturn(policy, "PolicyExists", false, nil)
				p = p.ApplyMethodReturn(policy, "CreateS3AccessPolicy", iam.GetMockTaskResponseUi(), nil)
				p = p.ApplyFuncReturn(iam.TaskStatus, true, nil)
				p = p.ApplyFuncReturn(iam.NewS3User, user)
				p = p.ApplyMethodReturn(user, "UserExists", false, nil)
				p = p.ApplyMethodReturn(user, "CreateS3User", "user_bucket1_user_key", nil)
				p = p.ApplyMethodReturn(user, "ApplyPolicy", iam.GetMockTaskResponseUi(), nil)
				return p
			},
			wantErr: false,
		},
		{
			name: "Failure due to missing secret",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: &bucketclientset.Clientset{},
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverGrantBucketAccessRequest{
					Name:               accessName,
					AuthenticationType: cosi.AuthenticationType_Key,
					Parameters: map[string]string{
						"cosiUserSecretNamespace": namespace,
						"cosiUserSecretName":      secretName,
					},
				},
			},
			want: &cosi.DriverGrantBucketAccessResponse{
				AccountId: accessName,
			},
			setupPatches: func() *gomonkey.Patches {
				p := gomonkey.ApplyFuncReturn(getSecret, nil, errors.New("failed to get secret"))
				return p
			},
			wantErr: true,
		},
		{
			name: "Failure due to missing parameter",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: &bucketclientset.Clientset{},
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverGrantBucketAccessRequest{
					Name:               accessName,
					AuthenticationType: cosi.AuthenticationType_Key,
				},
			},
			want: &cosi.DriverGrantBucketAccessResponse{
				AccountId: accessName,
			},
			setupPatches: func() *gomonkey.Patches {
				return nil
			},
			wantErr: true,
		},
		{
			name: "Failure due to missing GLCP User id",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: &bucketclientset.Clientset{},
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverGrantBucketAccessRequest{
					Name:               accessName,
					AuthenticationType: cosi.AuthenticationType_Key,
					Parameters: map[string]string{
						"cosiUserSecretNamespace": namespace,
						"cosiUserSecretName":      secretName,
					},
				},
			},
			want: &cosi.DriverGrantBucketAccessResponse{
				AccountId: accessName,
			},
			setupPatches: func() *gomonkey.Patches {
				p := gomonkey.ApplyFuncReturn(getSecret, removeSecretData(utils.GLCP_USER_CLIENTID), nil)
				p = p.ApplyFuncReturn(os.Getenv, "DUMMY_PROXY_URL")
				return p
			},
			wantErr: true,
		},
		{
			name: "Failure due to missing GLCP User secretkey",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: &bucketclientset.Clientset{},
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverGrantBucketAccessRequest{
					Name:               accessName,
					AuthenticationType: cosi.AuthenticationType_Key,
					Parameters: map[string]string{
						"cosiUserSecretNamespace": namespace,
						"cosiUserSecretName":      secretName,
					},
				},
			},
			want: &cosi.DriverGrantBucketAccessResponse{
				AccountId: accessName,
			},
			setupPatches: func() *gomonkey.Patches {
				p := gomonkey.ApplyFuncReturn(getSecret, removeSecretData(utils.GLCP_USER_SECRET_KEY), nil)
				p = p.ApplyFuncReturn(os.Getenv, "DUMMY_PROXY_URL")
				return p
			},
			wantErr: true,
		},
		{
			name: "Failure due to missing Dscc zone",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: &bucketclientset.Clientset{},
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverGrantBucketAccessRequest{
					Name:               accessName,
					AuthenticationType: cosi.AuthenticationType_Key,
					Parameters: map[string]string{
						"cosiUserSecretNamespace": namespace,
						"cosiUserSecretName":      secretName,
					},
				},
			},
			want: &cosi.DriverGrantBucketAccessResponse{
				AccountId: accessName,
			},
			setupPatches: func() *gomonkey.Patches {
				p := gomonkey.ApplyFuncReturn(getSecret, removeSecretData(utils.DSCC_ZONE), nil)
				p = p.ApplyFuncReturn(os.Getenv, "DUMMY_PROXY_URL")
				return p
			},
			wantErr: true,
		},
		{
			name: "Failure due to missing Cluster Serial Number",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: &bucketclientset.Clientset{},
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverGrantBucketAccessRequest{
					Name:               accessName,
					AuthenticationType: cosi.AuthenticationType_Key,
					Parameters: map[string]string{
						"cosiUserSecretNamespace": namespace,
						"cosiUserSecretName":      secretName,
					},
				},
			},
			want: &cosi.DriverGrantBucketAccessResponse{
				AccountId: accessName,
			},
			setupPatches: func() *gomonkey.Patches {
				p := gomonkey.ApplyFuncReturn(getSecret, removeSecretData(utils.ALLETRA_MP_X10K_SNO), nil)
				p = p.ApplyFuncReturn(os.Getenv, "DUMMY_PROXY_URL")
				return p
			},
			wantErr: true,
		},
		{
			name: "Failure due to missing proxy/glcp cloud URL in env",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: &bucketclientset.Clientset{},
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverGrantBucketAccessRequest{
					Name:               accessName,
					AuthenticationType: cosi.AuthenticationType_Key,
					Parameters: map[string]string{
						"cosiUserSecretNamespace": namespace,
						"cosiUserSecretName":      secretName,
					},
				},
			},
			want: &cosi.DriverGrantBucketAccessResponse{
				AccountId: accessName,
			},
			setupPatches: func() *gomonkey.Patches {
				p := gomonkey.ApplyFuncReturn(getSecret, secret, nil)
				return p
			},
			wantErr: true,
		},
		{
			name: "Failure due to error fetching access token",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: &bucketclientset.Clientset{},
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverGrantBucketAccessRequest{
					Name:               accessName,
					AuthenticationType: cosi.AuthenticationType_Key,
					Parameters: map[string]string{
						"cosiUserSecretNamespace": namespace,
						"cosiUserSecretName":      secretName,
					},
				},
			},
			want: &cosi.DriverGrantBucketAccessResponse{
				AccountId: accessName,
			},
			setupPatches: func() *gomonkey.Patches {
				p := gomonkey.ApplyFuncReturn(getSecret, secret, nil)
				p = p.ApplyFuncReturn(os.Getenv, "DUMMY_PROXY_URL")
				p = p.ApplyFuncReturn(getAccessToken, "", errors.New("failed to fetch access token"))
				return p
			},
			wantErr: true,
		},
		{
			name: "Failure due to empty access token",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: &bucketclientset.Clientset{},
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverGrantBucketAccessRequest{
					Name:               accessName,
					AuthenticationType: cosi.AuthenticationType_Key,
					Parameters: map[string]string{
						"cosiUserSecretNamespace": namespace,
						"cosiUserSecretName":      secretName,
					},
				},
			},
			want: &cosi.DriverGrantBucketAccessResponse{
				AccountId: accessName,
			},
			setupPatches: func() *gomonkey.Patches {
				p := gomonkey.ApplyFuncReturn(getSecret, secret, nil)
				p = p.ApplyFuncReturn(os.Getenv, "DUMMY_PROXY_URL")
				p = p.ApplyFuncReturn(getAccessToken, "", nil)
				return p
			},
			wantErr: true,
		},
		{
			name: "Failure due to error creating DSCC API client",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: &bucketclientset.Clientset{},
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverGrantBucketAccessRequest{
					Name:               accessName,
					AuthenticationType: cosi.AuthenticationType_Key,
					Parameters: map[string]string{
						"cosiUserSecretNamespace": namespace,
						"cosiUserSecretName":      secretName,
					},
				},
			},
			want: &cosi.DriverGrantBucketAccessResponse{
				AccountId: accessName,
			},
			setupPatches: func() *gomonkey.Patches {
				apiClient := iam.NewAPIClient("us1.xxxx.xxxxx.hpe.com", "dummytoken", "", "")
				p := gomonkey.ApplyFuncReturn(getSecret, secret, nil)
				p = p.ApplyFuncReturn(os.Getenv, "DUMMY_PROXY_URL")
				p = p.ApplyFuncReturn(getAccessToken, "bearerdummytoken", nil)
				p = p.ApplyFuncReturn(iam.NewAPIClient, apiClient)
				p = p.ApplyMethodReturn(apiClient, "GetAPIClient", nil, errors.New("failed to create API client"))
				return p
			},
			wantErr: true,
		},
		{
			name: "Failure due to error creating DSCC Task API client",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: &bucketclientset.Clientset{},
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverGrantBucketAccessRequest{
					Name:               accessName,
					AuthenticationType: cosi.AuthenticationType_Key,
					Parameters: map[string]string{
						"cosiUserSecretNamespace": namespace,
						"cosiUserSecretName":      secretName,
					},
				},
			},
			want: &cosi.DriverGrantBucketAccessResponse{
				AccountId: accessName,
			},
			setupPatches: func() *gomonkey.Patches {
				apiClient := iam.NewAPIClient("us1.xxxx.xxxxx.hpe.com", "dummytoken", "", "")
				p := gomonkey.ApplyFuncReturn(getSecret, secret, nil)
				p = p.ApplyFuncReturn(os.Getenv, "DUMMY_PROXY_URL")
				p = p.ApplyFuncReturn(getAccessToken, "bearerdummytoken", nil)
				p = p.ApplyFuncReturn(iam.NewAPIClient, apiClient)
				p = p.ApplyMethodReturn(apiClient, "GetAPIClient", nil, nil)
				p = p.ApplyMethodReturn(apiClient, "GetTaskAPIClient", nil, errors.New("failed to create Task API client"))
				return p
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				log:             tt.fields.log,
				K8sClientset:    tt.fields.K8sClientset,
				BucketClientset: tt.fields.BucketClientset,
			}
			patches := tt.setupPatches()
			if patches != nil {
				defer patches.Reset()
				patches.ApplyFunc(os.Exit, func(code int) {
					panic("os.Exit was called")
				})
			}
			got, err := s.DriverGrantBucketAccess(tt.args.ctx, tt.args.req)
			s.log.Error(err, "test error")
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.DriverCreateBucket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.DriverCreateBucket() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestDriverRevokeBucketAccess tests the DriverRevokeBucketAccess method for revoking access to a bucket.
// It ensures idempotency and validates the access revocation logic.
func TestDriverRevokeBucketAccess(t *testing.T) {
	type fields struct {
		log             logr.Logger
		K8sClientset    *kubernetes.Clientset
		BucketClientset bucketclientset.Interface
	}
	type args struct {
		ctx context.Context
		req *cosi.DriverRevokeBucketAccessRequest
	}

	// Create patches to retrieve secret
	systemId := "DUMMY_SERIAL_NUMBER"
	glcpCreds := getIAMCredentials()

	removeSecretData := func(remField string) *v1.Secret {
		data := map[string][]byte{utils.GLCP_USER_CLIENTID: []byte(glcpCreds.GLCPUser),
			utils.GLCP_USER_SECRET_KEY: []byte(glcpCreds.GLCPUserSecretKey),
			utils.DSCC_ZONE:            []byte(glcpCreds.DSCCZone),
			utils.ALLETRA_MP_X10K_SNO:  []byte(glcpCreds.SystemId),
		}
		delete(data, remField)
		return &v1.Secret{Data: data}
	}

	secret := createSecret(secretName, namespace, []byte("testuser"), []byte("testuser"), []byte("http://1.1.1.1:90"), &glcpCreds)
	coreV1 := &corev1.CoreV1Client{}
	client := &kubernetes.Clientset{}
	bucket := &v1alpha1.Bucket{
		ObjectMeta: metav1.ObjectMeta{
			Name: "bucket1",
		},
		Spec: v1alpha1.BucketSpec{
			Parameters: map[string]string{
				"cosiUserSecretNamespace": namespace,
				"cosiUserSecretName":      secretName,
			},
		},
	}
	bcs := fakebucketclientset.NewSimpleClientset(bucket)

	mockSecret := test_utils.MockSecret{Secret: secret}
	var secretInt corev1.SecretInterface = mockSecret
	patches := gomonkey.ApplyMethodReturn(client, "CoreV1", coreV1)
	defer patches.Reset()
	patches.ApplyMethodReturn(coreV1, "Secrets", secretInt)
	log := stdr.New(stdlog.New(os.Stdout, "", stdlog.LstdFlags))
	accessName := "bucket1_user"
	bucketName := "bucket1"
	policy := iam.NewS3Policy(utils.ACCESS_POLICY_PREFIX+accessName, bucketName, systemId, nil, context.Background())
	user := iam.NewS3User(utils.USER_PREFIX+accessName, systemId, nil, context.Background())

	tests := []struct {
		name         string
		fields       fields
		args         args
		want         *cosi.DriverRevokeBucketAccessResponse
		setupPatches func() *gomonkey.Patches
		wantErr      bool
	}{
		{
			name: "Revoke bucket access successfully",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: bcs,
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverRevokeBucketAccessRequest{
					AccountId: accessName,
					BucketId:  bucketName,
				},
			},
			want: &cosi.DriverRevokeBucketAccessResponse{},
			setupPatches: func() *gomonkey.Patches {
				p := gomonkey.ApplyFuncReturn(getSecret, secret, nil)
				p = p.ApplyFuncReturn(os.Getenv, "DUMMY_PROXY_URL")
				p = p.ApplyFuncReturn(getAccessToken, "bearerdummyoxyzxxzzz12xxxx341111zzzzyyyyyyQQQQQHHHHH", nil)
				p = p.ApplyFuncReturn(iam.NewS3Policy, policy)
				p = p.ApplyMethodReturn(policy, "PolicyExists", true, nil)
				p = p.ApplyMethodReturn(policy, "DeleteS3AccessPolicy", iam.GetMockTaskResponseUi(), nil)
				p = p.ApplyFuncReturn(iam.TaskStatus, true, nil)
				p = p.ApplyFuncReturn(iam.NewS3User, user)
				p = p.ApplyMethodReturn(user, "UserExists", true, nil)
				p = p.ApplyMethodReturn(user, "DeleteS3User", iam.GetMockTaskResponseUi(), nil)
				return p
			},
			wantErr: false,
		},
		{
			name: "Failure due to missing secret",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: bcs,
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverRevokeBucketAccessRequest{
					AccountId: accessName,
					BucketId:  bucketName,
				},
			},
			want: &cosi.DriverRevokeBucketAccessResponse{},
			setupPatches: func() *gomonkey.Patches {
				p := gomonkey.ApplyFuncReturn(getSecret, nil, errors.New("failed to get secret"))
				return p
			},
			wantErr: true,
		},
		{
			name: "Failure due to missing GLCP User id",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: bcs,
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverRevokeBucketAccessRequest{
					AccountId: accessName,
					BucketId:  bucketName,
				},
			},
			want: &cosi.DriverRevokeBucketAccessResponse{},
			setupPatches: func() *gomonkey.Patches {
				p := gomonkey.ApplyFuncReturn(getSecret, removeSecretData(utils.GLCP_USER_CLIENTID), nil)
				p = p.ApplyFuncReturn(os.Getenv, "DUMMY_PROXY_URL")
				return p
			},
			wantErr: true,
		},
		{
			name: "Failure due to missing GLCP User secretkey",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: bcs,
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverRevokeBucketAccessRequest{
					AccountId: accessName,
					BucketId:  bucketName,
				},
			},
			want: &cosi.DriverRevokeBucketAccessResponse{},
			setupPatches: func() *gomonkey.Patches {
				p := gomonkey.ApplyFuncReturn(getSecret, removeSecretData(utils.GLCP_USER_SECRET_KEY), nil)
				p = p.ApplyFuncReturn(os.Getenv, "DUMMY_PROXY_URL")
				return p
			},
			wantErr: true,
		},
		{
			name: "Failure due to missing Dscc zone",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: bcs,
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverRevokeBucketAccessRequest{
					AccountId: accessName,
					BucketId:  bucketName,
				},
			},
			want: &cosi.DriverRevokeBucketAccessResponse{},
			setupPatches: func() *gomonkey.Patches {
				p := gomonkey.ApplyFuncReturn(getSecret, removeSecretData(utils.DSCC_ZONE), nil)
				p = p.ApplyFuncReturn(os.Getenv, "DUMMY_PROXY_URL")
				return p
			},
			wantErr: true,
		},
		{
			name: "Failure due to missing Cluster Serial Number",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: bcs,
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverRevokeBucketAccessRequest{
					AccountId: accessName,
					BucketId:  bucketName,
				},
			},
			want: &cosi.DriverRevokeBucketAccessResponse{},
			setupPatches: func() *gomonkey.Patches {
				p := gomonkey.ApplyFuncReturn(getSecret, removeSecretData(utils.ALLETRA_MP_X10K_SNO), nil)
				p = p.ApplyFuncReturn(os.Getenv, "DUMMY_PROXY_URL")
				return p
			},
			wantErr: true,
		},
		{
			name: "Failure due to missing proxy/glcp cloud URL in env",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: bcs,
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverRevokeBucketAccessRequest{
					AccountId: accessName,
					BucketId:  bucketName,
				},
			},
			want: &cosi.DriverRevokeBucketAccessResponse{},
			setupPatches: func() *gomonkey.Patches {
				p := gomonkey.ApplyFuncReturn(getSecret, secret, nil)
				return p
			},
			wantErr: true,
		},
		{
			name: "Failure due to error fetching access token",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: bcs,
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverRevokeBucketAccessRequest{
					AccountId: accessName,
					BucketId:  bucketName,
				},
			},
			want: &cosi.DriverRevokeBucketAccessResponse{},
			setupPatches: func() *gomonkey.Patches {
				p := gomonkey.ApplyFuncReturn(getSecret, secret, nil)
				p = p.ApplyFuncReturn(os.Getenv, "DUMMY_PROXY_URL")
				p = p.ApplyFuncReturn(getAccessToken, "", errors.New("failed to fetch access token"))
				return p
			},
			wantErr: true,
		},
		{
			name: "Failure due to empty access token",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: bcs,
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverRevokeBucketAccessRequest{
					AccountId: accessName,
					BucketId:  bucketName,
				},
			},
			want: &cosi.DriverRevokeBucketAccessResponse{},
			setupPatches: func() *gomonkey.Patches {
				p := gomonkey.ApplyFuncReturn(getSecret, secret, nil)
				p = p.ApplyFuncReturn(os.Getenv, "DUMMY_PROXY_URL")
				p = p.ApplyFuncReturn(getAccessToken, "", nil)
				return p
			},
			wantErr: true,
		},
		{
			name: "Failure due to error creating DSCC API client",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: bcs,
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverRevokeBucketAccessRequest{
					AccountId: accessName,
					BucketId:  bucketName,
				},
			},
			want: &cosi.DriverRevokeBucketAccessResponse{},
			setupPatches: func() *gomonkey.Patches {
				apiClient := iam.NewAPIClient("us1.xxxx.xxxxx.hpe.com", "dummytoken", "", "")
				p := gomonkey.ApplyFuncReturn(getSecret, secret, nil)
				p = p.ApplyFuncReturn(os.Getenv, "DUMMY_PROXY_URL")
				p = p.ApplyFuncReturn(getAccessToken, "bearerdummytoken", nil)
				p = p.ApplyFuncReturn(iam.NewAPIClient, apiClient)
				p = p.ApplyMethodReturn(apiClient, "GetAPIClient", nil, errors.New("failed to create API client"))
				return p
			},
			wantErr: true,
		},
		{
			name: "Failure due to error creating DSCC Task API client",
			fields: fields{
				log:             log,
				K8sClientset:    &kubernetes.Clientset{},
				BucketClientset: bcs,
			},
			args: args{
				ctx: context.TODO(),
				req: &cosi.DriverRevokeBucketAccessRequest{
					AccountId: accessName,
					BucketId:  bucketName,
				},
			},
			want: &cosi.DriverRevokeBucketAccessResponse{},
			setupPatches: func() *gomonkey.Patches {
				apiClient := iam.NewAPIClient("us1.xxxx.xxxxx.hpe.com", "dummytoken", "", "")
				p := gomonkey.ApplyFuncReturn(getSecret, secret, nil)
				p = p.ApplyFuncReturn(os.Getenv, "DUMMY_PROXY_URL")
				p = p.ApplyFuncReturn(getAccessToken, "bearerdummytoken", nil)
				p = p.ApplyFuncReturn(iam.NewAPIClient, apiClient)
				p = p.ApplyMethodReturn(apiClient, "GetAPIClient", nil, nil)
				p = p.ApplyMethodReturn(apiClient, "GetTaskAPIClient", nil, errors.New("failed to create Task API client"))
				return p
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				log:             tt.fields.log,
				K8sClientset:    tt.fields.K8sClientset,
				BucketClientset: tt.fields.BucketClientset,
			}
			patches := tt.setupPatches()
			if patches != nil {
				defer patches.Reset()
				patches.ApplyFunc(os.Exit, func(code int) {
					panic("os.Exit was called")
				})
			}
			got, err := s.DriverRevokeBucketAccess(tt.args.ctx, tt.args.req)
			s.log.Error(err, "test error")
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.DriverCreateBucket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.DriverCreateBucket() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestCreateS3Client verifies the creation of an S3Client instance.
// It ensures that the client is initialized with the correct parameters.
func TestCreateS3Client(t *testing.T) {
	// ...existing code...
}

// TestParseBucketParams exercises every field-level and cross-field
// validation rule consumed by DriverCreateBucket. It is a pure unit test
// (no mocks, no I/O) because parseBucketParams is a pure function — that
// is the point of having factored it out.
func TestParseBucketParams(t *testing.T) {
	tests := []struct {
		name      string
		params    map[string]string
		wantErr   bool
		errSubstr string                                    // optional: assert the message names the offending rule
		check     func(t *testing.T, r utils.BucketRequest) // optional: assert produced struct
	}{
		// --- happy paths ----------------------------------------------------
		{
			name:   "all params omitted yields defaults",
			params: map[string]string{},
			check: func(t *testing.T, r utils.BucketRequest) {
				if r.Versioning != utils.FeatureDisabled {
					t.Errorf("versioning default = %q, want Disabled", r.Versioning)
				}
				if r.Compression != "" || r.Locking != "" || r.RetentionMode != "" ||
					r.ObjectLockDays != 0 || r.ObjectLockYears != 0 {
					t.Errorf("non-versioning fields should be zero, got %+v", r)
				}
			},
		},
		{
			name: "case-insensitive keys accepted",
			params: map[string]string{
				"Versioning":  "Enabled",
				"COMPRESSION": "Disabled",
			},
			check: func(t *testing.T, r utils.BucketRequest) {
				if r.Versioning != utils.FeatureEnabled {
					t.Errorf("versioning = %q, want Enabled", r.Versioning)
				}
				if r.Compression != utils.FeatureDisabled {
					t.Errorf("compression = %q, want Disabled", r.Compression)
				}
			},
		},
		{
			name: "full valid object-lock chain",
			params: map[string]string{
				"versioning":               "Enabled",
				"locking":                  "Enabled",
				"retentionMode":            "GOVERNANCE",
				"defaultRetentionInterval": "30d",
			},
			check: func(t *testing.T, r utils.BucketRequest) {
				if r.Locking != utils.FeatureEnabled ||
					r.RetentionMode != utils.RetentionModeGovernance ||
					r.ObjectLockDays != 30 || r.ObjectLockYears != 0 {
					t.Errorf("unexpected result: %+v", r)
				}
			},
		},
		{
			name: "years interval maps to ObjectLockYears",
			params: map[string]string{
				"versioning":               "Enabled",
				"locking":                  "Enabled",
				"retentionMode":            "COMPLIANCE",
				"defaultRetentionInterval": "2y",
			},
			check: func(t *testing.T, r utils.BucketRequest) {
				if r.ObjectLockDays != 0 || r.ObjectLockYears != 2 {
					t.Errorf("days/years = (%d,%d), want (0,2)", r.ObjectLockDays, r.ObjectLockYears)
				}
			},
		},
		{
			name: "months interval is normalised to days",
			params: map[string]string{
				"versioning":               "Enabled",
				"locking":                  "Enabled",
				"retentionMode":            "GOVERNANCE",
				"defaultRetentionInterval": "3m",
			},
			check: func(t *testing.T, r utils.BucketRequest) {
				if r.ObjectLockDays != 90 || r.ObjectLockYears != 0 {
					t.Errorf("days/years = (%d,%d), want (90,0)", r.ObjectLockDays, r.ObjectLockYears)
				}
			},
		},

		// --- field-level validation ----------------------------------------
		{
			name:      "invalid compression rejected",
			params:    map[string]string{"compression": "On"},
			wantErr:   true,
			errSubstr: `compression="On" is invalid`,
		},
		{
			name:      "invalid versioning rejected",
			params:    map[string]string{"versioning": "yes"},
			wantErr:   true,
			errSubstr: `versioning="yes" is invalid`,
		},
		{
			name:      "invalid locking rejected",
			params:    map[string]string{"locking": "maybe"},
			wantErr:   true,
			errSubstr: `locking="maybe" is invalid`,
		},
		{
			name:      "invalid retentionMode rejected",
			params:    map[string]string{"retentionMode": "WEAK"},
			wantErr:   true,
			errSubstr: `retentionMode="WEAK" is invalid`,
		},
		{
			name:      "invalid defaultRetentionInterval rejected",
			params:    map[string]string{"defaultRetentionInterval": "30q"},
			wantErr:   true,
			errSubstr: `defaultRetentionInterval="30q" is invalid`,
		},

		// --- cross-field validation (the new rules) ------------------------
		{
			name: "locking=Enabled without versioning rejected",
			params: map[string]string{
				"locking":                  "Enabled",
				"versioning":               "Disabled",
				"retentionMode":            "GOVERNANCE",
				"defaultRetentionInterval": "30d",
			},
			wantErr:   true,
			errSubstr: "versioning",
		},
		{
			name: "locking=Enabled with versioning omitted (defaults to Disabled) rejected",
			params: map[string]string{
				"locking":                  "Enabled",
				"retentionMode":            "GOVERNANCE",
				"defaultRetentionInterval": "30d",
			},
			wantErr:   true,
			errSubstr: "versioning",
		},
		{
			name: "retentionMode without locking rejected",
			params: map[string]string{
				"retentionMode": "GOVERNANCE",
			},
			wantErr:   true,
			errSubstr: "locking",
		},
		{
			name: "defaultRetentionInterval without locking rejected",
			params: map[string]string{
				"defaultRetentionInterval": "30d",
			},
			wantErr:   true,
			errSubstr: "locking",
		},
		{
			name: "retention settings with locking=Disabled rejected",
			params: map[string]string{
				"locking":                  "Disabled",
				"retentionMode":            "COMPLIANCE",
				"defaultRetentionInterval": "1y",
			},
			wantErr:   true,
			errSubstr: "locking",
		},
		{
			name: "locking=Enabled without retentionMode rejected",
			params: map[string]string{
				"versioning":               "Enabled",
				"locking":                  "Enabled",
				"defaultRetentionInterval": "30d",
			},
			wantErr:   true,
			errSubstr: "retentionMode",
		},
		{
			name: "locking=Enabled without defaultRetentionInterval rejected",
			params: map[string]string{
				"versioning":    "Enabled",
				"locking":       "Enabled",
				"retentionMode": "GOVERNANCE",
			},
			wantErr:   true,
			errSubstr: "defaultRetentionInterval",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseBucketParams(tt.params)
			if (err != nil) != tt.wantErr {
				t.Fatalf("parseBucketParams() err = %v, wantErr=%v", err, tt.wantErr)
			}
			if tt.wantErr {
				if tt.errSubstr != "" && err != nil && !strings.Contains(err.Error(), tt.errSubstr) {
					t.Errorf("error %q does not contain %q", err.Error(), tt.errSubstr)
				}
				return
			}
			if tt.check != nil {
				tt.check(t, got)
			}
		})
	}
}

// TestDriverCreateBucket_InvalidArgument confirms that every validation
// failure inside parseBucketParams surfaces to the gRPC client as
// codes.InvalidArgument. It is deliberately decoupled from the heavy mock
// chain in TestDriverCreateBucket: parameter validation rejects the request
// before CreateBucket / SetObjectLockConfiguration are ever called, so only
// createS3Client needs patching.
func TestDriverCreateBucket_InvalidArgument(t *testing.T) {
	mockS3c := &S3Client{
		BaseURL:    "http://mock-s3",
		HTTPClient: &http.Client{},
		AccessKey:  "k",
		SecretKey:  "k",
	}
	patches := gomonkey.ApplyFunc(createS3Client,
		func(_ context.Context, _ map[string]string, _ *kubernetes.Clientset) (*S3Client, error) {
			return mockS3c, nil
		})
	defer patches.Reset()

	log := stdr.New(stdlog.New(os.Stdout, "", stdlog.LstdFlags))
	s := &Server{
		log:             log,
		K8sClientset:    &kubernetes.Clientset{},
		BucketClientset: &bucketclientset.Clientset{},
	}

	cases := []struct {
		name   string
		params map[string]string
	}{
		{"bad feature value", map[string]string{"compression": "On"}},
		{"bad retentionMode", map[string]string{"retentionMode": "WEAK"}},
		{"bad interval", map[string]string{"defaultRetentionInterval": "30q"}},
		{"locking without versioning", map[string]string{
			"locking": "Enabled", "versioning": "Disabled",
			"retentionMode": "GOVERNANCE", "defaultRetentionInterval": "30d",
		}},
		{"retention without locking", map[string]string{
			"retentionMode": "GOVERNANCE", "defaultRetentionInterval": "30d",
		}},
		{"locking without retention settings", map[string]string{
			"versioning": "Enabled", "locking": "Enabled",
		}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := s.DriverCreateBucket(context.TODO(), &cosi.DriverCreateBucketRequest{
				Name:       "b",
				Parameters: c.params,
			})
			if err == nil {
				t.Fatalf("expected error, got nil")
			}
			if got := status.Code(err); got != codes.InvalidArgument {
				t.Errorf("status code = %v, want InvalidArgument; err=%v", got, err)
			}
		})
	}
}

func createSecret(secretName string, namespace string, accessKey []byte, secretKey []byte, endpoint []byte, glcpCreds *utils.IAMCredentials) *v1.Secret {
	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: namespace,
		},
		Type: v1.SecretTypeOpaque,
		Data: map[string][]byte{
			"accessKey": accessKey,
			"secretKey": secretKey,
			"endpoint":  endpoint,
		},
	}
	if glcpCreds != nil {
		secret.Data[utils.GLCP_USER_CLIENTID] = []byte(glcpCreds.GLCPUser)
		secret.Data[utils.GLCP_USER_SECRET_KEY] = []byte(glcpCreds.GLCPUserSecretKey)
		secret.Data[utils.DSCC_ZONE] = []byte(glcpCreds.DSCCZone)
		secret.Data[utils.ALLETRA_MP_X10K_SNO] = []byte(glcpCreds.SystemId)
	}
	return secret
}

func getIAMCredentials() utils.IAMCredentials {
	return utils.IAMCredentials{
		GLCPUser:          "xxxxxxx-xxx-123-3456",
		GLCPUserSecretKey: "zzzzzrandomxxxxxxzzzzz",
		DSCCZone:          "us1.xxxx.xxxxx.hpe.com",
		SystemId:          "DUMMY_SERIAL_NUMBER",
	}
}
