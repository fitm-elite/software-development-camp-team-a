package sheet

import (
	"encoding/csv"
	"errors"
	"os"

	"github.com/rs/zerolog/log"
)

// Errors of the sheet package.
var (
	ErrInvalidFileExtension = errors.New("invalid file extension")
)

// Sheet represents a sheet.
type Sheet struct {
	file *os.File
}

// Properties represents the properties of a sheet.
type Properties struct {
	path string
}

// OptionFunc represents an option function.
type OptionFunc func(properties *Properties)

// WithPath sets the path for the sheet properties.
func WithPath(path string) OptionFunc {
	return func(properties *Properties) {
		properties.path = path
	}
}

// New creates a new sheet.
func New(options ...OptionFunc) (*Sheet, error) {
	properties := &Properties{}
	for _, option := range options {
		option(properties)
	}

	file, err := os.Open(properties.path)
	if err != nil {
		return nil, err
	}

	sheet := &Sheet{
		file: file,
	}

	return sheet, nil
}

// Read reads the sheet.
func (s *Sheet) Read() ([][]string, error) {
	reader := csv.NewReader(s.file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Error().Err(err).Msg("failed to read CSV file")
		return nil, err
	}

	return records, nil
}

// Close closes the sheet.
func (s *Sheet) Close() error {
	return s.file.Close()
}
