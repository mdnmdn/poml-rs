# POML JS/TS API

This document describes the JS/TS API for the POML template engine, based on the official documentation.

The JS/TS SDK allows you to work with POML using a JSX-like syntax in your TypeScript or JavaScript code.

## Installation

```bash
npm install pomljs
```

## Core Functions

The JS/TS API is centered around two core functions: `read` and `write`.

- **`read(prompt)`:** Parses the prompt components (written in JSX-like syntax) into an intermediate representation (IR).
- **`write(ir)`:** Renders the IR into different formats (e.g., markdown).

## Quick Start

Here is a quick example of how to use the JS/TS API:

```typescript
import { Paragraph, Image } from 'poml/essentials';
import { read, write } from 'poml';

const prompt = <Paragraph>
  Hello, world! Here is an image:
  <Image src="photo.jpg" alt="A beautiful scenery" />
</Paragraph>;

// Parse the prompt components into an intermediate representation (IR)
const ir = await read(prompt);

// Render it to markdown
const markdown = write(ir);
```

## Modules

The JS/TS SDK is organized into several modules:

- **`essentials`:** Provides the basic POML components like `<Paragraph>` and `<Image>`.
- **`base`:** Contains the base components and utilities.
- **`cli`:** The command-line interface.
- **`file`:** Utilities for working with files.
- **`writer`:** The rendering engine that writes the IR to different formats.
