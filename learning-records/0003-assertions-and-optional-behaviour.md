# 0003 — Assertions are for asking what is in the box, not for reshaping it

- **Date:** 2026-09-03
- **Chapter:** 7 — Types, Methods, and Interfaces
- **Status:** taught, exercise outstanding
- **Follows:** [0002](0002-shallow-copies-and-method-sets.md)

## Context

Closes a thread opened in [0001](0001-adapters-over-widening-types.md): Brad asked whether a Go
conversion is "basically a cast," as in `double(1)`. Answer: yes for the C sense, but Java/C#
"cast" bundles two operations that Go deliberately separates.

## What was learned

1. Conversion `T(x)` is compile-time and cannot fail. Assertion `x.(T)` is runtime and can.
2. An assertion's operand must be an **interface value** — there has to be a box to open.
   `l.(Ranker)` on a `*League` is a compile error, not a runtime one.
3. Comma-ok is the default form; the naked form panics, though with a precise message
   (`interface conversion: main.Ranker is main.RankerAdapter, not *main.League`).
4. Type switches narrow the bound variable only in single-type cases. In `default`, or a
   multi-type case, it keeps the interface type.
5. `errors.As` is the assertion you will write most, and it must be `errors.As` rather than a
   hand-rolled assertion, because `%w`-wrapped errors nest.

## The non-obvious bit

The **optional-behaviour probe**: accept a narrow interface, then assert at runtime for a wider
one, with a fallback. This is not a workaround — it is how the standard library adds capability
without breaking callers, and Brad has been benefiting from it unknowingly since exercise 3.
`RankPrinter` calls `io.WriteString`, which probes for `io.StringWriter`; `os.Stdout` satisfies
it, so those calls have been taking the allocation-free path all along.

Verified: `os.Stdout`, `*bytes.Buffer`, and `*strings.Builder` all satisfy `io.StringWriter`.

## Consequence

Interfaces stop being only a constraint to satisfy and become a runtime question you can ask.
This is the hinge for the next lesson, "accept interfaces, return structs" — a concrete return
type is what lets callers probe for more later.

## Open threads

- Exercise not yet done: add `WinCount` to `*League` and the probe to `RankPrinter`; confirm
  `RankerStruct` prints counts and `RankerAdapter` does not.
- `RankPrinter` builds its error with `%v`, discarding the underlying cause. Should be `%w`.
- `errcheck`: first `RankPrinter` call site still discards its error (carried from 0002).
- Lesson 4 candidate: accept interfaces, return structs.
