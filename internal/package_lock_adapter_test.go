package internal

import (
	"errors"
	"os"
	"strings"
	"testing"
)

func TestDiffWithNoDependencies(t *testing.T) {
	adapter := PackageLockAdapter{}

	file1 := strings.NewReader("{}")
	file2 := strings.NewReader("{}")

	diff, err := adapter.GetDiff(file1, file2)

	if err != nil {
		t.Errorf("Expected error to be nil, got: %+v", err)
	}

	if len(diff) > 0 {
		t.Errorf("Expected diff to be empty, got: %+v", diff)
	}
}

func TestDiffParsingNoJsonFile(t *testing.T) {
	adapter := PackageLockAdapter{}

	file1 := strings.NewReader("")
	file2 := strings.NewReader("")

	_, err := adapter.GetDiff(file1, file2)
	if !errors.Is(err, JSON_FILE_ERROR) {
		t.Errorf("Expected error to %+v, got: %+v", JSON_FILE_ERROR, err)
	}
}

func TestDiffReturnDifferences(t *testing.T) {
	adapter := PackageLockAdapter{}

	file1, _ := os.Open("../internal/test_files/package-lock-app-v0.json")
	file2, _ := os.Open("../internal/test_files/package-lock-app-v1.json")

	diff, err := adapter.GetDiff(file1, file2)
	if err != nil {
		t.Errorf("Expected error to be nil, got: %+v", err)
	}

	expected := 23

	if len(diff) != expected {
		t.Errorf("Expected diff to have %d elements, got: %d", expected, len(diff))
	}

	file3, _ := os.Open("../internal/test_files/package-lock-axios1.json")
	file4, _ := os.Open("../internal/test_files/package-lock-axios2.json")

	diff2, err := adapter.GetDiff(file3, file4)
	if err != nil {
		t.Errorf("Expected error to be nil, got: %+v", err)
	}

	expected = 1

	if len(diff2) != expected {
		t.Errorf("Expected diff to have %d elements, got: %d", expected, len(diff2))
	}
}

func TestDiffReturnDifferentFileOrder(t *testing.T) {
	adapter := PackageLockAdapter{}

	file1, _ := os.Open("../internal/test_files/package-lock-app-v0.json")
	file2, _ := os.Open("../internal/test_files/package-lock-app-v1.json")

	diff, err := adapter.GetDiff(file2, file1)
	if err != nil {
		t.Errorf("Expected error to be nil, got: %+v", err)
	}

	expected := 23

	if len(diff) != expected {
		t.Errorf("Expected diff to have %d elements, got: %d", expected, len(diff))
	}
}

func TestDiffReturnDifferentPackageFormatFiles(t *testing.T) {
	adapter := PackageLockAdapter{}

	file1, _ := os.Open("../internal/test_files/package-lock-npm6.json")
	file2, _ := os.Open("../internal/test_files/package-lock-app-v1.json")

	diff, err := adapter.GetDiff(file2, file1)
	if err != nil {
		t.Errorf("Expected error to be nil, got: %+v", err)
	}

	expected := 18

	if len(diff) != expected {
		t.Errorf("Expected diff to have %d elements, got: %d", expected, len(diff))
	}
}
