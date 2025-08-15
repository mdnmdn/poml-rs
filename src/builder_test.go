package poml

import (
	"reflect"
	"testing"
)

func TestBuilder_Simple(t *testing.T) {
	b := NewBuilder()
	b.Add("p", map[string]string{"speaker": "human"})
	b.Text("Hello")
	b.End()

	// Inspect the tree directly
	if len(b.root.Children) != 1 {
		t.Fatalf("Expected 1 child for root, got %d", len(b.root.Children))
	}
	p, ok := b.root.Children[0].(*Element)
	if !ok {
		t.Fatal("Expected child to be an element.")
	}
	if p.Tag != "p" || p.Attr["speaker"] != "human" {
		t.Errorf("Unexpected element data. Got tag=%s, attrs=%v", p.Tag, p.Attr)
	}
	if len(p.Children) != 1 {
		t.Fatalf("Expected 1 child for <p>, got %d", len(p.Children))
	}
	text, ok := p.Children[0].(*Text)
	if !ok || text.Content != "Hello" {
		t.Errorf("Unexpected text content. Got: %v", p.Children[0])
	}
}

func TestBuilder_Nested(t *testing.T) {
	b := NewBuilder()
	b.Add("div", nil).
		Add("p", nil).
		Text("Inner").
		End(). // End p
		Add("p", nil).
		Text("Another inner").
		End(). // End p
		End()  // End div

	// Manually inspect the built tree
	expected := &Element{
		Tag:  "poml",
		Attr: map[string]string{},
		Children: []Node{
			&Element{
				Tag:  "div",
				Attr: map[string]string{},
				Children: []Node{
					&Element{
						Tag:  "p",
						Attr: map[string]string{},
						Children: []Node{
							&Text{Content: "Inner"},
						},
					},
					&Element{
						Tag:  "p",
						Attr: map[string]string{},
						Children: []Node{
							&Text{Content: "Another inner"},
						},
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(b.root, expected) {
		t.Errorf("Built tree does not match expected tree.\nGot:      %#v\nExpected: %#v", b.root, expected)
	}
}

func TestBuilder_Render(t *testing.T) {
	b := NewBuilder()
	b.Add("p", nil).Text("Hello, {{ name }}").End()

	context := map[string]interface{}{"name": "Builder"}
	result, err := b.Render(context)
	if err != nil {
		t.Fatalf("Builder.Render() failed: %v", err)
	}

	expected := "Hello, Builder"
	if result != expected {
		t.Errorf("Expected render output '%s', got '%s'", expected, result)
	}
}
