---
status: proposed
---
# Immutable Tokens

## Context and Problem Statement

While hardening the lexer logic, I wanted to consolidate token logic/helpers into the token package, and was also curious how things would be different if the tokens were immutable.

## Decision Outcome

- ctors
  - `token.New()`
  - `token.NewIdent()`
- accessors
  - `Token.Literal()`
  - `Token.Type()`
- helpers
  - `Token.Is()`
