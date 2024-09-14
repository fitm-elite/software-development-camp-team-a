package command

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/fitm-elite/elebs/packages/sheet"
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

			for idx, record := range records {
				if idx == 0 {
					continue
				}

				fmt.Println(record)
			}

			// reader := file.Read()
			// for {
			// 	reader.ReadAll()

			// 	record, err := reader.Read()
			// 	if err == io.EOF {
			// 		break
			// 	}
			// 	if err != nil {
			// 		log.Error().Err(err).Msg("failed to read CSV file")
			// 		break
			// 	}
			// 	fmt.Println(record)
			// }
		},
	}

	root.AddCommand(use)

	return root.Execute()
}
