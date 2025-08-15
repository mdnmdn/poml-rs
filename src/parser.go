package poml

import (
	"bytes"
	"encoding/xml"
	"io"
)

// Node represents a node in the POML document tree.
// It can be either an Element or a Text node.
type Node interface {
	isNode()
}

// Element represents a POML element, like <p> or <task>.
type Element struct {
	Tag      string
	Attr     map[string]string
	Children []Node
}

// Text represents a raw text node within an element.
type Text struct {
	Content string
}

// isNode() ensures that only Element and Text can be a Node.
func (e Element) isNode() {}
func (t Text) isNode()    {}

// Parse reads from an io.Reader and returns the root Element of the POML document.
// This initial version focuses on parsing the XML structure without handling templating.
func Parse(r io.Reader) (*Element, error) {
	decoder := xml.NewDecoder(r)

	// Find the root element
	var root *Element
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if se, ok := token.(xml.StartElement); ok {
			root = &Element{
				Tag:  se.Name.Local,
				Attr: make(map[string]string),
			}
			for _, attr := range se.Attr {
				root.Attr[attr.Name.Local] = attr.Value
			}
			root.Children, err = parseChildren(decoder)
			if err != nil {
				return nil, err
			}
			break // Found root, break out
		}
	}

	if root == nil {
		return nil, io.EOF // Or a more specific error
	}

	return root, nil
}

func parseChildren(decoder *xml.Decoder) ([]Node, error) {
	var children []Node
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		switch t := token.(type) {
		case xml.StartElement:
			element := &Element{
				Tag:  t.Name.Local,
				Attr: make(map[string]string),
			}
			for _, attr := range t.Attr {
				element.Attr[attr.Name.Local] = attr.Value
			}
			element.Children, err = parseChildren(decoder)
			if err != nil {
				return nil, err
			}
			children = append(children, element)
		case xml.CharData:
			// Ignore whitespace-only text nodes between elements
			if len(bytes.TrimSpace(t)) > 0 {
				children = append(children, &Text{Content: string(t)})
			}
		case xml.EndElement:
			return children, nil
		}
	}
	return children, nil
}
