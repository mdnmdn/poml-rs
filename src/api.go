package poml

import (
	"os"
	"strings"
)

// RenderFromString parses a POML string and renders it with the given context.
// This is the primary entry point for using the POML engine.
func RenderFromString(pomlString string, context map[string]interface{}) (string, error) {
	reader := strings.NewReader(pomlString)
	root, err := Parse(reader)
	if err != nil {
		return "", err
	}
	return render(root, context)
}

// RenderFromFile reads a POML file and renders it with the given context.
func RenderFromFile(filepath string, context map[string]interface{}) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	root, err := Parse(file)
	if err != nil {
		return "", err
	}
	return render(root, context)
}
