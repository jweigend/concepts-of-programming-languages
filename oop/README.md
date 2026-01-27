# Object-Oriented Programming

## Purpose

This directory contains examples demonstrating **object-oriented programming concepts** in Go.

## Subdirectories

| Directory | Topic | Description |
|-----------|-------|-------------|
| `boolparser/` | Parser | Lexer/Parser/AST pattern |
| `mail/` | Interfaces | Dependency injection |
| `polymorphism/` | Polymorphism | Interfaces vs embedding |
| `rational/` | Value Types | Custom numeric type |
| `stack/` | Data Structures | Generic container |

## Key Concepts

- **Interfaces** - Implicit implementation (duck typing)
- **Struct embedding** - Composition over inheritance
- **Method receivers** - Attaching behavior to types
- **Encapsulation** - Unexported fields (lowercase)

## Go vs Traditional OOP

| Concept | Go | Java/C++ |
|---------|-----|----------|
| Inheritance | Embedding | Class hierarchy |
| Polymorphism | Interfaces | Virtual methods |
| Encapsulation | Package-level | Access modifiers |
| Constructors | Factory functions | Constructor methods |

## Run

```bash
go test ./...
```
