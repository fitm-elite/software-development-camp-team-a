package command

import "testing"

func TestFileExtensionValidator_Success(t *testing.T) {
	t.Parallel()

	ok := fileExtensionValidator([]string{"example.csv"})
	if !ok {
		t.Fatal("Expected true, got false")
	}
}

func TestFileExtensionValidator_Fail(t *testing.T) {
	t.Parallel()

	ok := fileExtensionValidator([]string{"example.txt"})
	if ok {
		t.Fatal("Expected false, got true")
	}
}
