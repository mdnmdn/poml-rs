# POML Rust SDK Implementation Plan

This document outlines the plan for creating the Rust SDK for POML.

## Development Flow

For each phase and task outlined below, the following process will be followed:
1. At the beginning of the task, run `cargo build` and `cargo test` to ensure the codebase is in a clean state.
2. Implement the feature or task.
3. At the end of the task, run `cargo build`, `cargo test`, and `cargo fmt` to ensure the changes are high quality and have not introduced regressions.
4. Upon successful completion, this plan file will be updated to mark the task as complete.

## Implementation Phases

### Phase 1: Project Setup & Core Parsing
- [ ] **Initialize Project Structure**: Set up the Rust project with `cargo`, creating the `src`, `tests`, and `examples` directories.
- [ ] **Implement POML Parser**: Develop a parser that converts POML markup into a Rust struct tree. A library like `quick-xml` or `xml-rs` could be used.

### Phase 2: Template Engine Implementation
- [ ] **Expression Evaluation**: Implement evaluation of `{{...}}` JavaScript expressions. A library like `deno_core` or `rquickjs` could be used.
- [ ] **Implement `if` attribute**: Implement conditional rendering based on the `if` attribute.
- [ ] **Implement `for` attribute**: Implement list-based iteration using the `for` attribute, including the `loop` variable.
- [ ] **Implement `<let>` tag**: Implement variable definition from values, inline JSON, and external files.
- [ ] **Implement `<include>` tag**: To allow including content from other `.poml` files.
- [ ] **Implement `<stylesheet>` tag**: To handle CSS-like styling rules.

### Phase 3: Idiomatic Rust SDK API
- [ ] **Core `render` Function**: Design and implement the primary `poml::render()` function, with variants like `render_from_string` and `render_from_file`.
- [ ] **Programmatic Builder API**: Design and implement an idiomatic Rust builder API (e.g., using the builder pattern) for creating POML documents programmatically.

### Phase 4: Comprehensive Testing
- [ ] **Port Original Tests**: Adapt the official test suites from the Python/JS SDKs to ensure full compatibility.

### Phase 5: Documentation and Finalization
- [ ] **Write Rust Examples**: Populate the `examples` directory with Rust programs demonstrating common use cases.
- [ ] **Document the Rust API**: Write clear, comprehensive Rust doc comments (`///`) and potentially a `_docs/rust-api.md` file.
- [ ] **Update `agents.md`**: Ensure `_docs/agents.md` is up-to-date.
