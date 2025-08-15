package poml

import (
	"os"
	"testing"
)

func TestRenderFromString(t *testing.T) {
	pomlString := `<poml>Hello, {{ name }}</poml>`
	context := map[string]interface{}{"name": "API"}

	result, err := RenderFromString(pomlString, context)
	if err != nil {
		t.Fatalf("RenderFromString() failed: %v", err)
	}

	expected := "Hello, API"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestRenderFromFile(t *testing.T) {
	// Create a temp file
	pomlContent := `<poml>Hello from a file, {{ name }}</poml>`
	tmpFile, err := os.CreateTemp("", "test-*.poml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	t.Cleanup(func() { os.Remove(tmpFile.Name()) })

	if _, err := tmpFile.Write([]byte(pomlContent)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	// Test RenderFromFile
	context := map[string]interface{}{"name": "World"}
	result, err := RenderFromFile(tmpFile.Name(), context)
	if err != nil {
		t.Fatalf("RenderFromFile() failed: %v", err)
	}

	expected := "Hello from a file, World"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
