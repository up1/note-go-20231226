package main

import (
	"os"
	"testing"
)

func TestMainFunction(t *testing.T) {
	// Create a temporary file
	tempFile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write some data to the file
	data := []byte("Hello, World!")
	if _, err := tempFile.Write(data); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	// Call the function with the temp file
	body, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	// Check the result
	if string(body) != string(data) {
		t.Fatalf("Expected %s, got %s", data, body)
	}
}
