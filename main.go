package main

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"

	"github.com/fitm-elite/elebs/command"
	localContext "github.com/fitm-elite/elebs/packages/context"
	"github.com/fitm-elite/elebs/packages/linebot"
	"github.com/fitm-elite/elebs/packages/logger"
	"github.com/fitm-elite/elebs/packages/minio"
	"github.com/fitm-elite/elebs/packages/timezone"
)

// main is the entry point of the application.
func main() {
	var err error
	ctx := context.Background()

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

	messagingApi, err := linebot.New(
		linebot.WithMessagingApi(),
	)
	if err != nil {
		log.Panic().Err(err).Msg("failed to create LineBot client")
	}

	ctx = context.WithValue(ctx, localContext.MinioKeyContextKey, minioClient)
	ctx = context.WithValue(ctx, localContext.MessagingApiContextKey, messagingApi)

	if err = command.Execute(ctx); err != nil {
		log.Panic().Err(err).Msg("failed to execute command")
	}
}
