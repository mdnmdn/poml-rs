package poml

import (
	"strings"
	"testing"
)

func TestStylesheetAppliesAttributes(t *testing.T) {
	pomlString := `
<poml>
  <stylesheet>
    {
      "p": { "syntax": "json" },
      ".special": { "class": "extra-special" }
    }
  </stylesheet>
  <p>Some text</p>
  <div class="special">Some div</div>
</poml>
`
	root, err := Parse(strings.NewReader(pomlString))
	if err != nil {
		t.Fatalf("Parse() failed: %v", err)
	}

	stylesheet, err := parseStylesheet(root)
	if err != nil {
		t.Fatalf("parseStylesheet() failed: %v", err)
	}

	if stylesheet == nil {
		t.Fatal("Stylesheet was not found or parsed.")
	}

	stylesheet.Apply(root)

	// After applying the stylesheet, inspect the tree
	var pElement, divElement *Element
	for _, node := range root.Children {
		if elem, ok := node.(*Element); ok {
			if elem.Tag == "p" {
				pElement = elem
			}
			if elem.Tag == "div" {
				divElement = elem
			}
		}
	}

	if pElement == nil {
		t.Fatal("Could not find the <p> element in the tree.")
	}
	if syntax, ok := pElement.Attr["syntax"]; !ok || syntax != "json" {
		t.Errorf("Expected 'p' tag to have attribute syntax='json', but it did not. Attrs: %v", pElement.Attr)
	}

	if divElement == nil {
		t.Fatal("Could not find the <div> element in the tree.")
	}
	if class, ok := divElement.Attr["class"]; !ok || class != "extra-special" {
		t.Errorf("Expected '.special' div to have attribute class='extra-special', but it did not. Attrs: %v", divElement.Attr)
	}
}
