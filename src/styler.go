package poml

import (
	"encoding/json"
	"strings"
)

// Stylesheet defines the structure for styling rules.
// The key is the selector (e.g., "p" or ".csv").
// The value is a map of attribute keys to attribute values.
type Stylesheet map[string]map[string]string

// Apply walks the node tree and applies the stylesheet rules.
func (s Stylesheet) Apply(node Node) {
	if node == nil {
		return
	}

	if element, ok := node.(*Element); ok {
		// Apply styles to this element
		for selector, styles := range s {
			if matches(selector, element) {
				for key, value := range styles {
					// Stylesheet attributes overwrite existing ones.
					element.Attr[key] = value
				}
			}
		}

		// Recursively apply to children
		for _, child := range element.Children {
			s.Apply(child)
		}
	}
}

// matches checks if an element matches a given selector.
func matches(selector string, element *Element) bool {
	if strings.HasPrefix(selector, ".") {
		// Class selector
		className := strings.TrimPrefix(selector, ".")
		return element.Attr["class"] == className
	}
	// Tag selector
	return element.Tag == selector
}

// parseStylesheet finds a stylesheet in the tree, parses it, and removes the node.
func parseStylesheet(root *Element) (Stylesheet, error) {
	var stylesheet Stylesheet
	var stylesheetNode *Element
	var nodeIndex int

	// Find the stylesheet node
	for i, node := range root.Children {
		if elem, ok := node.(*Element); ok && elem.Tag == "stylesheet" {
			stylesheetNode = elem
			nodeIndex = i
			break
		}
	}

	if stylesheetNode != nil {
		// Remove the stylesheet from the tree
		root.Children = append(root.Children[:nodeIndex], root.Children[nodeIndex+1:]...)

		// Parse the JSON content
		if len(stylesheetNode.Children) > 0 {
			if text, ok := stylesheetNode.Children[0].(*Text); ok {
				err := json.Unmarshal([]byte(text.Content), &stylesheet)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return stylesheet, nil
}
