package sheet_test

import (
	"testing"

	"github.com/fitm-elite/elebs/packages/sheet"
)

func TestNewSheet_Success(t *testing.T) {
	t.Parallel()

	_, err := sheet.New(sheet.WithPath("./../../example.csv"))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestNewSheet_InvalidFileExtension(t *testing.T) {
	t.Parallel()

	_, err := sheet.New(sheet.WithPath("./../../example.txt"))
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
}

func TestNewSheet_IsFiveColumns(t *testing.T) {
	t.Parallel()

	sheet, err := sheet.New(sheet.WithPath("./../../example.csv"))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	records, err := sheet.Read()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	for idx, record := range records {
		if idx == 0 {
			if len(record) != 5 {
				t.Fatalf("Expected 5 columns, got %v", len(record))
			}

			break
		}
	}
}

func TestNewSheet_EmptyProperties(t *testing.T) {
	t.Parallel()

	t.Run("Empty path", func(t *testing.T) {
		_, err := sheet.New()
		if err == nil {
			t.Fatal("Expected an error, got nil")
		}
	})
}
