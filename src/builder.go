package poml

// Builder provides a fluent API for programmatically constructing POML documents.
type Builder struct {
	root  *Element
	stack []*Element // The top of the stack is the current element being built.
}

// NewBuilder creates and returns a new Builder instance, initialized with a root <poml> element.
func NewBuilder() *Builder {
	root := &Element{Tag: "poml", Attr: make(map[string]string)}
	return &Builder{
		root:  root,
		stack: []*Element{root},
	}
}

// currentElement returns the element at the top of the stack.
func (b *Builder) currentElement() *Element {
	if len(b.stack) == 0 {
		return nil
	}
	return b.stack[len(b.stack)-1]
}

// Add adds a new element with the given tag and attributes as a child of the current element.
// It also pushes the new element onto the stack, making it the new current element.
func (b *Builder) Add(tag string, attrs map[string]string) *Builder {
	parent := b.currentElement()
	if parent == nil {
		// Should not happen if used correctly
		return b
	}

	if attrs == nil {
		attrs = make(map[string]string)
	}

	element := &Element{Tag: tag, Attr: attrs}
	parent.Children = append(parent.Children, element)
	b.stack = append(b.stack, element)
	return b
}

// Text adds a text node as a child of the current element.
func (b *Builder) Text(content string) *Builder {
	current := b.currentElement()
	if current == nil {
		return b
	}
	current.Children = append(current.Children, &Text{Content: content})
	return b
}

// End pops the current element off the stack, effectively moving up one level in the tree.
// It should be called to close the current element.
func (b *Builder) End() *Builder {
	if len(b.stack) > 1 { // Cannot pop the root element
		b.stack = b.stack[:len(b.stack)-1]
	}
	return b
}

// Render finalizes the POML tree and renders it with the given context.
func (b *Builder) Render(context map[string]interface{}) (string, error) {
	return render(b.root, context)
}
