package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// New returns a new zerolog logger that configured to write to stderr and include timestamp and caller information.
func NewZerolog() zerolog.Logger {
	return zerolog.New(
		func() io.Writer {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
			zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
			writer := io.Writer(
				zerolog.ConsoleWriter{
					Out:        os.Stderr,
					TimeFormat: time.RFC3339Nano,
				},
			)

			return writer
		}(),
	).With().
		Timestamp().
		Caller().
		Logger()
}
