# 0001 — Adapt at the boundary instead of widening the domain type

- **Date:** 2026-09-02
- **Chapter:** 7 — Types, Methods, and Interfaces
- **Status:** learned

## Context

`ch7/main.go` defines `type TeamName string` and `func (l League) Ranking() []TeamName`.
Exercise 3 introduces `Ranker`, an interface requiring `Ranking() []string`. `League` does
not satisfy it:

```
League does not implement Ranker (wrong type for method Ranking)
	have Ranking() []TeamName
	want Ranking() []string
```

The obvious fix is to delete `TeamName` and use `string` everywhere. Brad deliberately kept
the mismatch in order to write an adapter.

## What was learned

1. Interface satisfaction in Go is **exact**, not compatible. Same name, same parameters,
   same results, or it does not implement.
2. `[]TeamName` cannot be converted to `[]string`. Slice conversion requires *identical*
   element types, even though the element types themselves convert freely. The copy loop is
   unavoidable, and that loop is the substance of the adapter.
3. A defined type (`type B A`) starts with an **empty method set**. Embedding is what carries
   methods across.
4. A func type is a defined type, so it can have methods. That is the entire mechanism behind
   `http.HandlerFunc`, and it is the shape used here:
   `type RankerAdapter func() []string` plus
   `func (f RankerAdapter) Ranking() []string { return f() }`.
5. A **method value** (`league.RankingStrings`, no parentheses) is a function with the
   receiver pre-bound, so it converts directly into the adapter type.

## The non-obvious bit

Both first attempts failed the same way: a method was written whose *name* referred to the
adapter or the interface (`RankerAdapter()`, `Ranker()`) rather than being the method the
interface actually demands, on a type that will actually be passed. Interface satisfaction
never looks at names of types — only at the method set of the value handed over.

## Consequence for future lessons

Brad raises the difficulty of book exercises on purpose to hit a concept head-on. Never
resolve that difficulty by removing it. Explain the mechanism.

## Outcome

Both shapes were written by hand and both work (`go vet` clean, correct output). Brad chose
the func adapter as the better fit here and gave the right reason: the interface has one
method, so the struct wrapper's extra machinery buys nothing.

Two vocabulary corrections landed alongside it:

- A conversion (`RankerAdapter(f)`) is compile-time checked and always succeeds. A **type
  assertion** (`x.(T)`) is the runtime-checked one. "Cast" is not a Go term and blurs the two.
- The method is `Ranking`; `Ranker` is the interface. Confusing the two names is precisely
  what broke both earlier attempts.

## Open threads

- Pointer vs value method sets. `MatchResult` has a pointer receiver, `Ranking` does not, and
  `RankerStruct` embeds `League` **by value** — so a promoted `MatchResult` call through the
  wrapper mutates a copied struct whose `Wins` map is nonetheless shared. Posed as a
  prediction exercise; this is lesson 2.
- Error returns from `RankPrinter` are discarded at both call sites. `go vet` misses it,
  `errcheck` does not.
