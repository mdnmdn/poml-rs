package poml

import (
	"os"
	"strings"
	"testing"
)

func TestRenderIncludeTag(t *testing.T) {
	// Create a snippet file
	snippetContent := `<world>World</world>`
	snippetFile, err := os.CreateTemp("", "snippet-*.poml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	t.Cleanup(func() { os.Remove(snippetFile.Name()) })

	if _, err := snippetFile.Write([]byte(snippetContent)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	snippetFile.Close()

	// Create a main file that includes the snippet
	mainContent := `<poml>Hello, <include src="` + snippetFile.Name() + `" />!</poml>`

	// Parse and Render the main content
	root, err := Parse(strings.NewReader(mainContent))
	if err != nil {
		t.Fatalf("Parse() failed: %v", err)
	}

	result, err := render(root, nil)
	if err != nil {
		t.Fatalf("render() failed: %v", err)
	}

	expected := "Hello, World!"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
