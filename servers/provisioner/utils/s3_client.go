// © Copyright 2024 Hewlett Packard Enterprise Development LP

// Package utils
package utils

import (
	"context"
	"errors"
	stdlog "log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

const (
	// Bucket Class parameters' keys
	bucketTagsKey              string = "bucketTags"
	CosiUserSecretNameKey      string = "cosiUserSecretName"
	CosiUserSecretNamespaceKey string = "cosiUserSecretNamespace"

	// Provisioner pod's env variable for the COSI driver container
	PodNamespaceEnv string = "POD_NAMESPACE"
)

var (
	httpTimeout     time.Duration = 30 * time.Second
	enablePathStyle bool          = true
	disableSSL      bool          = true
)

var (
	errInvalidTagFormat = errors.New("invalid format used to specify tags in the Bucket parameters. Expected: key=value,key=value")
	errEmptyTagKey      = errors.New("tag key cannot be empty. Expected: key=value,key=value")
)

type S3Client struct {
	Client s3iface.S3API
}

type BucketParameters struct {
	BucketTags *s3.Tagging
}

// Logger key to extract logger from context
const LoggerKey = "loggerKey"

// Create a new S3 client
func NewS3Client(ctx context.Context, accessKey string, secretKey string, endpoint string) (*S3Client, error) {
	// Set log level to debug
	logLevel := aws.LogDebug

	// Create HTTP client
	client := &http.Client{
		Timeout: httpTimeout,
	}

	// Create a new session
	session, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:         aws.String(endpoint),
		Region:           aws.String(endpoints.UsEast1RegionID),
		DisableSSL:       aws.Bool(disableSSL),
		HTTPClient:       client,
		LogLevel:         aws.LogLevel(logLevel),
		MaxRetries:       aws.Int(3),
		S3ForcePathStyle: aws.Bool(enablePathStyle),
	})
	if err != nil {
		return nil, err
	}

	// Create a new s3 client
	svc := s3.New(session)
	return &S3Client{
		Client: svc,
	}, nil
}

// Uses the s3 client to create a bucket
func (s *S3Client) CreateBucket(ctx context.Context, bucketName string, param map[string]string) error {
	log := GetLoggerFromContext(ctx)

	log.Info("Attempting to create bucket", "bucketName", bucketName)
	_, err := s.Client.CreateBucket(&s3.CreateBucketInput{Bucket: aws.String(bucketName)})
	if err != nil {
		awsErr, ok := err.(awserr.Error)
		if ok {
			if awsErr.Code() == s3.ErrCodeBucketAlreadyExists {
				log.Info("The bucket already exists.", "bucketName", bucketName)
			} else if awsErr.Code() == s3.ErrCodeBucketAlreadyOwnedByYou {
				log.Info("The bucket is already owned by you.", "bucketName", bucketName)
			} else {
				log.Error(err, "failed to create bucket using S3 client")
				return err
			}
		} else {
			log.Error(err, "failed to create bucket using S3 client")
			return err
		}
	}

	log.Info("Checking if bucket tags have been defined", "bucketName", bucketName)
	if bp, err := getBucketTags(ctx, param); err != nil {
		e := s.DeleteBucket(ctx, bucketName)
		if e != nil {
			log.Error(e, "failed to clean up bucket after parsing bucket tags failed")
		}
		// return the error that caused parsing of bucket tags to fails
		return err
	} else if bp.BucketTags != nil && len(bp.BucketTags.TagSet) != 0 {
		log.Info("Attempting to set the tags for bucket", "tags", bp.BucketTags.TagSet, "bucketName", bucketName)
		// Call PutBucketTagging
		_, err = s.Client.PutBucketTagging(&s3.PutBucketTaggingInput{Bucket: aws.String(bucketName), Tagging: bp.BucketTags})
		if err != nil {
			log.Error(err, "failed to apply the tagging configuration to the bucket", "tags", bp.BucketTags.TagSet, "bucketName", bucketName)
			e := s.DeleteBucket(ctx, bucketName)
			if e != nil {
				log.Error(e, "failed to clean up bucket after application of tag configuration failed")
			}
			// return the error that caused putting bucket tags to fail
			return err
		}
	}

	return nil
}

// Uses the s3 client to delete a bucket
func (s *S3Client) DeleteBucket(ctx context.Context, bucketName string) error {
	log := GetLoggerFromContext(ctx)
	log.Info("Attempting to delete bucket", "bucketName", bucketName)
	_, err := s.Client.DeleteBucket(&s3.DeleteBucketInput{Bucket: aws.String(bucketName)})
	if err != nil {
		log.Error(err, "failed to delete bucket using S3 client")
	}

	return err
}

// Get the bucket tag parameter and parse if any bucket tags have been defined
func getBucketTags(ctx context.Context, param map[string]string) (*BucketParameters, error) {
	log := GetLoggerFromContext(ctx)
	bp := &BucketParameters{}

	// Get the value of the Bucket tag parameter key
	val := param[bucketTagsKey]

	// If the value is not empty, proceed to parse
	if val != "" {
		// convert the string of bucket tags to *s3.Tagging
		bt, err := parseBucketTags(ctx, val)
		if err != nil {
			return nil, err
		}
		bp.BucketTags = bt
	} else {
		log.Info("No bucket tags found")
		return bp, nil
	}

	return bp, nil
}

// parse the value of the bucketTags parameter (string) to an array of tags ([]*s3.Tag)
func parseBucketTags(ctx context.Context, paramValue string) (*s3.Tagging, error) {
	log := GetLoggerFromContext(ctx)
	bucketTags := &s3.Tagging{}
	// Split the parameter value into an array of key-value pairs using the ',' delimiter
	tags := strings.Split(paramValue, ",")

	// Iterate over each key-value pair
	for _, tag := range tags {
		// Ignore empty string tags and trailing commas
		if tag == "" {
			continue
		}

		// Split the tag into its component key and value using the '=' delimiter
		keyValue := strings.Split(tag, "=")

		// If the length of array 'keyValue' is not 2, return an error due to invalid format
		if len(keyValue) != 2 {
			log.Error(errInvalidTagFormat, "failed to parse the Bucket tags.",
				"bucketTag", keyValue)
			return nil, errInvalidTagFormat
		}

		// Check if the bucket tag key is empty. An empty bucket tag value is allowed.
		key := strings.TrimSpace(keyValue[0])
		if key != "" {
			// append the key-value pair to the TagSet array
			bucketTags.TagSet = append(bucketTags.TagSet, &s3.Tag{Key: aws.String(key), Value: aws.String(strings.TrimSpace(keyValue[1]))})
		} else {
			log.Error(errEmptyTagKey, "failed to parse Bucket tag due to invalid Key.",
				"receivedBucketTag", keyValue)
			return nil, errEmptyTagKey
		}
	}

	return bucketTags, nil
}

func GetLoggerFromContext(ctx context.Context) *logr.Logger {
	logger := ctx.Value(LoggerKey)

	// If the context has no logger, create a new logger and return
	if logger == nil {
		log := stdr.New(stdlog.New(os.Stdout, "", stdlog.LstdFlags))
		return &log
	}

	return logger.(*logr.Logger)
}
