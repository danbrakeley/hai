---
status: accepted
---
# Token Enum Type

## Context and Problem Statement

On page 16, strings are used instead of an integer type for tokens.

The book says this keeps code examples simple. I expect using integer token types will be an overall more performant choice, and know that with the help of generators, can also be easy to use, with type-specific helpers for parsing between ints and strings.

So not really a problem, so much as a personal challenge based on an assumption that may, in the end, turn out not to be true and/or worthwhile (who says it is more performant? w/out benchmarks comparing both options, it is unknown).

## Considered Options

* using strings as in the book
* using ints: [stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)
* using ints: [dmarkham/enumer](https://github.com/dmarkham/enumer)

## Decision Outcome

`dmarkham/enumer`:

### Consequences

Will need to translate any use of strings in the book text to using my chosen type and generated helpers.

## More Information

* Install: `go install github.com/dmarkham/enumer@latest`
  * See magefile's Setup target.
* Use: `//go:generate enumer -type=YOURTYPE -json`
