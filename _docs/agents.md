# POML Rust Porting Project

This document outlines the goals and structure of the Rust port for the POML template engine.

## Goals

The primary goal of this project is to create a Rust implementation of the POML template engine that is compatible with the existing Python and JS/TS implementations.

The Rust port should:
- Parse and render `.poml` files.
- Support the full POML language specification, including the template engine features.
- Provide a simple and intuitive API for developers.
- Be well-documented and easy to use.

## Project Structure

The project will be structured as follows:

- `_docs`: Contains the documentation for the project, including this file.
- `src`: Contains the Rust source code for the POML engine.
- `examples`: Will contain example `.poml` files and Rust programs demonstrating how to use the engine.
- `tests`: Integration tests will be located in the `tests` directory. Unit tests will be co-located with the code they are testing.

## References

- **Official POML Documentation:** [https://microsoft.github.io/poml/latest/](https://microsoft.github.io/poml/latest/)
- **Official POML Repository:** [https://github.com/microsoft/poml](https://github.com/microsoft/poml)
