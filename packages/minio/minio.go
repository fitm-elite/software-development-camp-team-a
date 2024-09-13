package minio

import (
	"errors"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// MinIO errors.
var (
	ErrEmptyProperties = errors.New("endpoint, accessKeyID, secretAccessKey cannot be empty")
)

// Properties represents a MinIO properties.
type Properties struct {
	endpoint        string
	accessKeyID     string
	secretAccessKey string
	useSSL          bool
}

// OptionFunc represents a function that configures a MinIO properties.
type OptionFunc func(properties *Properties)

// New creates a new MinIO properties.
func WithEndpoint(endpoint string) OptionFunc {
	return func(properties *Properties) {
		properties.endpoint = endpoint
	}
}

// WithAccessKeyID sets the access key ID for the MinIO properties.
func WithAccessKeyID(accessKeyID string) OptionFunc {
	return func(properties *Properties) {
		properties.accessKeyID = accessKeyID
	}
}

// WithSecretAccessKey sets the secret access key for the MinIO properties.
func WithSecretAccessKey(secretAccessKey string) OptionFunc {
	return func(properties *Properties) {
		properties.secretAccessKey = secretAccessKey
	}
}

// WithUseSSL sets whether the MinIO properties should use SSL.
func WithUseSSL(useSSL bool) OptionFunc {
	return func(properties *Properties) {
		properties.useSSL = useSSL
	}
}

// New creates a new MinIO client.
func New(options ...OptionFunc) (*minio.Client, error) {
	properties := &Properties{}
	for _, option := range options {
		option(properties)
	}

	if properties.endpoint == "" || properties.accessKeyID == "" || properties.secretAccessKey == "" {
		return nil, ErrEmptyProperties
	}

	minioClient, err := minio.New(
		properties.endpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(properties.accessKeyID, properties.secretAccessKey, ""),
			Secure: properties.useSSL,
		},
	)
	if err != nil {
		return nil, err
	}

	return minioClient, nil
}
