# POML Go SDK Implementation Plan

This document outlines the plan for creating the Go SDK for POML.

## Development Flow

For each phase and task outlined below, the following process will be followed:
1.  At the beginning of the task, run `go build`, `go test`, `go fmt`, and `go vet` to ensure the codebase is in a clean state.
2.  Implement the feature or task.
3.  At the end of the task, run `go build`, `go test`, `go fmt`, and `go vet` again to ensure the changes are high quality and have not introduced regressions.
4.  Upon successful completion, this plan file will be updated to mark the task as complete.

## Implementation Phases

### Phase 1: Project Setup & Core Parsing (Completed)
- [x] **Initialize Project Structure**: Created the `src`, `tests`, and `examples` directories and initialized a Go module.
- [x] **Implement POML Parser**: Developed a parser that converts POML markup into a Go struct tree.

### Phase 2: Template Engine Implementation (In Progress)
- [x] **Expression Evaluation**: Implemented evaluation of `{{...}}` JavaScript expressions using the `goja` library.
- [x] **Implement `if` attribute**: Implemented conditional rendering based on the `if` attribute.
- [x] **Implement `for` attribute**: Implemented list-based iteration using the `for` attribute, including the `loop` variable.
- [x] **Implement `<let>` tag**: Implemented variable definition from values, inline JSON, and external files.
- [x] **Implement `<include>` tag**: To allow including content from other `.poml` files.
- [x] **Implement `<stylesheet>` tag**: To handle CSS-like styling rules.

### Phase 3: Idiomatic Go SDK API (Completed)
- [x] **Core `Render` Function**: Finalized the primary `poml.Render()` function signature and behavior by creating `RenderFromString` and `RenderFromFile`.
- [x] **Programmatic Builder API**: Designed and implemented an idiomatic Go builder API for creating POML documents programmatically.

### Phase 4: Comprehensive Testing
- [ ] **Port Original Tests**: Adapt the official test suites from the Python/JS SDKs to ensure full compatibility.

### Phase 5: Documentation and Finalization
- [ ] **Write Go Examples**: Populate the `examples` directory with Go programs demonstrating common use cases.
- [ ] **Document the Go API**: Write clear, comprehensive Go doc comments and a `_docs/go-api.md` file.
- [ ] **Update `agents.md`**: Add a reference to this plan file in `_docs/agents.md`.
