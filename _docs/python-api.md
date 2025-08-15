# POML Python API

This document describes the Python API for the POML template engine, based on the official Python SDK.

## `poml()` function

The `poml()` function is the main entry point for the Python API. It processes POML markup and returns the result in the specified format.

```python
poml(
    markup: str | Path,
    context: dict | str | Path | None = None,
    stylesheet: dict | str | Path | None = None,
    chat: bool = True,
    output_file: str | Path | None = None,
    format: OutputFormat = "dict",
    extra_args: Optional[List[str]] = None,
) -> list | dict | str:
```

### Parameters

- `markup`: The POML markup to process, as a string or a file path.
- `context`: The context data to inject into the template.
- `stylesheet`: The stylesheet to use for rendering.
- `chat`: Whether to process as a chat conversation.
- `output_file`: The path to save the output to.
- `format`: The output format (`raw`, `dict`, `openai_chat`, `langchain`, `pydantic`).
- `extra_args`: Additional command-line arguments to pass to the POML processor.

## `Prompt` class

The `Prompt` class provides a Pythonic way to build POML markup programmatically using context managers.

```python
from poml import Prompt

with Prompt() as p:
    with p.paragraph():
        with p.task(id="task1", status="open"):
            p.text("This is a task description.")
        with p.paragraph():
            p.text("This is a paragraph in the document.")

# Get the rendered output
prompt_output = p.render()

# Get the generated XML for debugging
xml_output = p.dump_xml()
```

### Methods

- `render()`: Renders the final XML and returns the output.
- `dump_xml()`: Returns the generated XML string.
- `text()`: Adds text content to the currently open element.
- `tag()`: Creates a new tag. You can also use methods like `p.paragraph()` as a shortcut.
