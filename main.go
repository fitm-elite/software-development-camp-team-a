package main

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"

	"github.com/fitm-elite/elebs/command"
	"github.com/fitm-elite/elebs/packages/linebot"
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

	minioClient, err := minio.New(
		minio.WithEndpoint(os.Getenv("MINIO_ENDPOINT")),
		minio.WithAccessKeyID(os.Getenv("MINIO_ACCESS_KEY_ID")),
		minio.WithSecretAccessKey(os.Getenv("MINIO_SECRET_ACCESS_KEY")),
		minio.WithUseSSL(true),
	)
	if err != nil {
		log.Panic().Err(err).Msg("failed to create MinIO client")
	}

	messagingApi, messagingApiBlob, err := linebot.New(
		linebot.WithMessagingApi(), linebot.WithMessagingApiBlob(),
	)
	if err != nil {
		log.Panic().Err(err).Msg("failed to create LineBot client")
	}

	_ = minioClient
	_ = messagingApi
	_ = messagingApiBlob

	if err = command.Execute(); err != nil {
		log.Panic().Err(err).Msg("failed to execute command")
	}
}
