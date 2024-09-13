package main

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"

	"github.com/fitm-elite/elebs/command"
	"github.com/fitm-elite/elebs/packages/logger"
	"github.com/fitm-elite/elebs/packages/minio"
	"github.com/fitm-elite/elebs/packages/timezone"
)

// main is the entry point of the application.
func main() {
	var err error

	time.Local = timezone.NewAsiaBangkok()
	log.Logger = logger.NewZerolog()

	if err = godotenv.Load(); err != nil {
		log.Fatal().Err(err).Msg("failed to load environment variables")
	}

	minioEndpoint := os.Getenv("MINIO_ENDPOINT")
	minioAccessKeyID := os.Getenv("MINIO_ACCESS_KEY_ID")
	minioSecretAccessKey := os.Getenv("MINIO_SECRET_ACCESS_KEY")
	minioClient, err := minio.New(
		minio.WithEndpoint(minioEndpoint),
		minio.WithAccessKeyID(minioAccessKeyID),
		minio.WithSecretAccessKey(minioSecretAccessKey),
		minio.WithUseSSL(true),
	)
	if err != nil {
		log.Panic().Err(err).Msg("failed to create MinIO client")
	}

	_ = minioClient

	if err = command.Execute(); err != nil {
		log.Panic().Err(err).Msg("failed to execute command")
	}
}
