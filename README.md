# POML Go SDK

[![Go CI](https://github.com/YOUR_USERNAME/YOUR_REPONAME/actions/workflows/go-ci.yml/badge.svg)](https://github.com/YOUR_USERNAME/YOUR_REPONAME/actions/workflows/go-ci.yml)

This project is a Go implementation of the POML (Prompt Orchestration Markup Language) template engine. It provides a comprehensive, idiomatic Go SDK for parsing and rendering `.poml` files, including support for the full POML template language specification.

## Features

*   **POML Parser**: A robust parser for the POML syntax.
*   **Template Engine**: Full support for template features, including:
    *   `{{...}}` JavaScript expression evaluation.
    *   `if` attribute for conditional rendering.
    *   `for` attribute for list iteration.
    *   `<let>` tag for defining variables from values, JSON, or files.
    *   `<include>` tag for embedding other POML files.
    *   `<stylesheet>` tag for applying attributes to elements.
*   **Idiomatic Go API**:
    *   Simple, user-friendly functions like `RenderFromString` and `RenderFromFile`.
    *   A fluent `Builder` API for programmatically constructing POML documents.

## Getting Started

*TODO: Add installation and usage examples here.*

---
**Note**: Please replace `YOUR_USERNAME/YOUR_REPONAME` in the CI badge URL with the actual path to your repository to activate the status badge.