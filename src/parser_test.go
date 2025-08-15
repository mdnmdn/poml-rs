package poml

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	pomlString := `<poml>
		<p speaker="human">Hello, world!</p>
	</poml>`

	reader := strings.NewReader(pomlString)
	root, err := Parse(reader)
	if err != nil {
		t.Fatalf("Parse() failed with error: %v", err)
	}

	if root.Tag != "poml" {
		t.Errorf("Expected root tag 'poml', got '%s'", root.Tag)
	}

	if len(root.Children) != 1 {
		t.Fatalf("Expected 1 child for root, got %d", len(root.Children))
	}

	// Type assertion to get the paragraph element
	pElement, ok := root.Children[0].(*Element)
	if !ok {
		t.Fatalf("Expected child to be an Element, but it was not.")
	}

	if pElement.Tag != "p" {
		t.Errorf("Expected element tag 'p', got '%s'", pElement.Tag)
	}

	if speaker, ok := pElement.Attr["speaker"]; !ok || speaker != "human" {
		t.Errorf("Expected attribute speaker='human', got '%s'", speaker)
	}

	if len(pElement.Children) != 1 {
		t.Fatalf("Expected 1 child for <p> element, got %d", len(pElement.Children))
	}

	// Type assertion to get the text node
	textNode, ok := pElement.Children[0].(*Text)
	if !ok {
		t.Fatalf("Expected child of <p> to be a Text node, but it was not.")
	}

	// The parser will pick up the whitespace and newlines as text nodes.
	// We need to trim the space to get the actual content.
	trimmedContent := strings.TrimSpace(textNode.Content)
	if trimmedContent != "Hello, world!" {
		t.Errorf("Expected text content 'Hello, world!', got '%s'", trimmedContent)
	}
}
