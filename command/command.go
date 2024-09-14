package command

import (
	"context"
	"fmt"
  
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog/log"
	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"

	localContext "github.com/fitm-elite/elebs/packages/context"
	"github.com/fitm-elite/elebs/packages/promptpay"
	"github.com/fitm-elite/elebs/packages/sheet"
	"github.com/fitm-elite/elebs/packages/utility"
)

// root is the root command object.
var root = &cobra.Command{
	Use:   "elebs",
	Short: "Elebs is a CLI tool for scrape data from a csv file that containing electric bills data.",
	Long:  "Elebs is a command-line tool designed to scrape and extract data from CSV files containing electric bills for KMUTNB Dormitory, Prachinburi campus. The tool automates the process of reading, filtering, and parsing electric bill data, Ideal for dormitory management and tenants, Elebs enables quick access to utility data and simplifies the workflow for handling multiple records at once. Its focus is on accuracy, speed, and ease of use for administrative tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World!")
	},
}

// Execute runs the root command.
func Execute(ctx context.Context) error {
	use := &cobra.Command{
		Use:   "use [file.csv]",
		Short: "Use a csv file to scrape data",
		Long:  "Use a csv file to scrape data for calculating the electric bill and push message to linebot.",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if ok := fileExtensionValidator(args); !ok {
				log.Fatal().Err(sheet.ErrInvalidFileExtension).Msg("invalid file extension")
			}

			file, err := sheet.New(sheet.WithPath(args[0]))
			if err != nil {
				log.Fatal().Err(err).Msg("failed to open file")
			}
			defer func() {
				if err := file.Close(); err != nil {
					log.Fatal().Err(err).Msg("failed to close file")
				}
			}()

			records, err := file.Read()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to read CSV file")
			}

			minioClient, ok := ctx.Value(localContext.MinioKeyContextKey).(*minio.Client)
			if !ok {
				log.Error().Err(err).Msg("failed to get minio client")
				return
			}

			messagingApi, ok := ctx.Value(localContext.MessagingApiContextKey).(*messaging_api.MessagingApiAPI)
			if !ok {
				log.Error().Err(err).Msg("failed to get messaging api")
				return
			}

			var validateWg sync.WaitGroup
			for _, record := range records[1:] {
				validateWg.Add(1)

				go func(record []string) {
					defer validateWg.Done()

					if record[0] == "" || record[1] == "" || record[2] == "" || record[3] == "" || record[4] == "" {
						log.Fatal().Msg("record cannot be empty")
					}

					roomNumber, err := strconv.ParseInt(record[0], 10, 64)
					if err != nil {
						log.Fatal().Err(err).Msg("failed to parse int")
					}

					if roomNumber < 200 || roomNumber > 527 {
						log.Fatal().Msg("room number is invalid")
					}

					lineIds := strings.Split(record[1], ",")
					if len(lineIds) == 0 {
						log.Fatal().Msg("line ids are empty")
						return
					}

					residents, err := strconv.ParseUint(record[4], 10, 64)
					if err != nil {
						log.Fatal().Err(err).Msg("failed to parse uint")
						return
					}

					if len(lineIds) != int(residents) {
						log.Fatal().Msg("line ids and residents are not equal")
						return
					}
				}(record)
			}

			validateWg.Wait()

			var recordWg sync.WaitGroup
			for _, record := range records[1:] {
				recordWg.Add(1)

				go func(record []string) {
					defer recordWg.Done()

					lineIds := strings.Split(record[1], ",")
					if len(lineIds) == 0 {
						log.Fatal().Msg("line ids are empty")
						return
					}

					residents, err := strconv.ParseUint(record[4], 10, 64)
					if err != nil {
						log.Fatal().Err(err).Msg("failed to parse uint")
						return
					}

					var rwg sync.WaitGroup
					for _, lineId := range lineIds {
						rwg.Add(1)

						go func(lineId *string) {
							defer rwg.Done()

							profile, err := messagingApi.GetProfile(*lineId)
							if err != nil {
								log.Error().Err(err).Msg("failed to get profile")
								return
							}

							billCost, err := strconv.ParseFloat(record[3], 64)
							if err != nil {
								log.Error().Err(err).Msg("failed to parse float")
								return
							}

							promptPayId := "0641823735"
							costDivided := utility.CostDivider(billCost, int(residents))

							promptpay := promptpay.PromptPay{
								PromptPayID: promptPayId,
								Amount:      costDivided,
							}
							promptPayCrc, err := promptpay.Gen()
							if err != nil {
								log.Error().Err(err).Msg("failed to generate promptpay")
								return
							}

							if err = qrcode.WriteFile(promptPayCrc, qrcode.Medium, 256, fmt.Sprintf("qrcode-%s.png", promptPayCrc)); err != nil {
								log.Error().Err(err).Msg("failed to write file")
								return
							}

							if _, err = minioClient.FPutObject(
								ctx, "elebs",
								fmt.Sprintf("qrcode-%s.png", promptPayCrc),
								fmt.Sprintf("qrcode-%s.png", promptPayCrc),
								minio.PutObjectOptions{},
							); err != nil {
								log.Error().Err(err).Msg("failed to put object")
								return
							}

							if err := os.Remove(fmt.Sprintf("qrcode-%s.png", promptPayCrc)); err != nil {
								log.Error().Err(err).Msg("failed to remove file")
								return
							}

							urlParams := make(url.Values)
							urlParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename=\"%s\"", fmt.Sprintf("qrcode-%s.png", promptPayCrc)))
							presigned, err := minioClient.PresignedGetObject(ctx, "elebs", fmt.Sprintf("qrcode-%s.png", promptPayCrc), 7*time.Hour, urlParams)
							if err != nil {
								log.Error().Err(err).Msg("failed to presigned object")
								return
							}

							if _, err = messagingApi.PushMessage(
								&messaging_api.PushMessageRequest{
									To: profile.UserId,
									Messages: []messaging_api.MessageInterface{
										&messaging_api.TextMessage{
											Text: fmt.Sprintf("ค่าไฟฟ้าของคุณ %s คิดเป็น %.2f บาท (จากทั้งหมด %.2f บาท) ณ ของวันที่ %v *โปรดชำระภายในวันที่ 25 ของทุกเดือน*", profile.DisplayName, costDivided, billCost, record[2]),
										},
									},
								}, "",
							); err != nil {
								log.Error().Err(err).Msg("failed to push message")
								return
							}

							if _, err = messagingApi.PushMessage(
								&messaging_api.PushMessageRequest{
									To: profile.UserId,
									Messages: []messaging_api.MessageInterface{
										&messaging_api.ImageMessage{
											OriginalContentUrl: presigned.String(),
											PreviewImageUrl:    presigned.String(),
										},
									},
								},
								"",
							); err != nil {
								log.Error().Err(err).Msg("failed to push message")
								return
							}

						}(&lineId)
					}

					rwg.Wait()
				}(record)
			}

			recordWg.Wait()
		},
	}

	root.AddCommand(use)

	return root.Execute()
}
