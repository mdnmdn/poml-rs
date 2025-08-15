# POML Language Specification

This document describes the specification of the POML language, based on the official documentation.

## Basic Syntax

POML uses an XML-like syntax with tags, attributes, and content.

- **Tags:** Mark the beginning and end of an element (e.g., `<p>`, `</p>`).
- **Attributes:** Provide additional information about an element (e.g., `<p speaker="human">`).
- **Content:** The text or other elements between the opening and closing tags.

The root of a POML document should be a `<poml>` tag.

## Template Engine

POML includes a powerful template engine with the following features:

### Expressions

Expressions are enclosed in double curly brackets (`{{ }}`) and can be used to evaluate variables and JavaScript expressions.

```poml
<poml>
  <p>Hello, {{name}}!</p>
</poml>
```

### Let Expressions

The `<let>` tag is used to define variables and import data.

- **Set a variable from a value:**
  ```poml
  <let name="greeting" value="Hello, world!" />
  ```
- **Import data from a file:**
  ```poml
  <let name="users" src="users.json" />
  ```
- **Set a variable from inline JSON:**
  ```poml
  <let name="person">
    {
      "name": "Alice",
      "age": 30
    }
  </let>
  ```

### For Attribute

The `for` attribute is used to loop over a list.

```poml
<poml>
  <list>
    <item for="item in ['apple', 'banana', 'cherry']">{{item}}</item>
  </list>
</poml>
```

Loop variables (`loop.index`, `loop.length`, `loop.first`, `loop.last`) are available inside the loop.

### If Condition

The `if` attribute is used for conditional rendering.

```poml
<poml>
  <p if="isVisible">This paragraph is visible.</p>
</poml>
```

### Include Files

The `<include>` tag is used to include content from another file.

```poml
<poml>
  <include src="snippet.poml" />
</poml>
```

## Stylesheet

The `<stylesheet>` tag allows you to apply CSS-like styles (attributes) to elements.

```poml
<poml>
  <stylesheet>
    {
      "p": {
        "syntax": "json"
      }
    }
  </stylesheet>
  <p>This text will be rendered as JSON.</p>
</poml>
```

Elements can be targeted by tag name or by `className` (e.g., `.csv`).

## Escape Characters

POML uses the following escape characters:

- `#quot;` for `"`
- `#apos;` for `'`
- `#amp;` for `&`
- `#lt;` for `<`
- `#gt;` for `>`
- `#hash;` for `#`
- `#lbrace;` for `{`
- `#rbrace;` for `}`
