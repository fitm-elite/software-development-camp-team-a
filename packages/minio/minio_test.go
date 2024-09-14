package minio_test

import (
	"testing"

	"github.com/fitm-elite/elebs/packages/minio"
)

// TestNewMinioClient_Success tests the successful creation of a MinIO client.
func TestNewMinioClient_Success(t *testing.T) {
	t.Parallel()

	client, err := minio.New(
		minio.WithEndpoint("localhost:9000"),
		minio.WithAccessKeyID("testAccessKey"),
		minio.WithSecretAccessKey("testSecretKey"),
		minio.WithUseSSL(false),
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if client == nil {
		t.Fatal("Expected a valid MinIO client, got nil")
	}
}

// TestNewMinioClient_EmptyProperties tests for error when required properties are missing.
func TestNewMinioClient_EmptyProperties(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		options []minio.OptionFunc
	}{
		{
			name: "Empty endpoint",
			options: []minio.OptionFunc{
				minio.WithAccessKeyID("testAccessKey"),
				minio.WithSecretAccessKey("testSecretKey"),
				minio.WithUseSSL(false),
			},
		},
		{
			name: "Empty accessKeyID",
			options: []minio.OptionFunc{
				minio.WithEndpoint("localhost:9000"),
				minio.WithSecretAccessKey("testSecretKey"),
				minio.WithUseSSL(false),
			},
		},
		{
			name: "Empty secretAccessKey",
			options: []minio.OptionFunc{
				minio.WithEndpoint("localhost:9000"),
				minio.WithAccessKeyID("testAccessKey"),
				minio.WithUseSSL(false),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := minio.New(tt.options...)

			if client != nil {
				t.Fatal("Expected nil client, got valid client")
			}
			if err == nil {
				t.Fatal("Expected an error, got nil")
			}
			if err != minio.ErrEmptyProperties {
				t.Fatalf("Expected ErrEmptyProperties, got %v", err)
			}
		})
	}
}

// TestNewMinioClient_ErrorOnConnection tests failure when MinIO client cannot connect.
func TestNewMinioClient_ErrorOnConnection(t *testing.T) {
	client, err := minio.New(
		minio.WithEndpoint("invalid-endpoint"),
		minio.WithAccessKeyID("testAccessKey"),
		minio.WithSecretAccessKey("testSecretKey"),
		minio.WithUseSSL(false),
	)
	if err != nil {
		t.Fatal("Expected error due to invalid endpoint, got none")
	}

	if client == nil {
		t.Fatal("Expected nil client, got valid client")
	}
}

// TestNewMinioClient_WithSSL tests MinIO client creation with SSL.
func TestNewMinioClient_WithSSL(t *testing.T) {
	client, err := minio.New(
		minio.WithEndpoint("localhost:9000"),
		minio.WithAccessKeyID("testAccessKey"),
		minio.WithSecretAccessKey("testSecretKey"),
		minio.WithUseSSL(true),
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if client == nil {
		t.Fatal("Expected a valid MinIO client, got nil")
	}
}
