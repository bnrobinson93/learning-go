# Resources

## Primary

- **[The Go Programming Language Specification](https://go.dev/ref/spec)** — the
  authoritative source. Short enough to read end to end. Sections used so far:
  - [Method sets](https://go.dev/ref/spec#Method_sets) — what determines whether a
    type implements an interface.
  - [Conversions](https://go.dev/ref/spec#Conversions) — why `[]TeamName` cannot be
    converted to `[]string`.
  - [Struct types](https://go.dev/ref/spec#Struct_types) — embedding and field/method
    promotion.
  - [Interface types](https://go.dev/ref/spec#Interface_types).
  - [Method values](https://go.dev/ref/spec#Method_values) — `v.M` with the receiver pre-bound.
  - [Address operators](https://go.dev/ref/spec#Address_operators) — the definitive list of
    what is addressable.
  - [Type assertions](https://go.dev/ref/spec#Type_assertions) and
    [Type switches](https://go.dev/ref/spec#Type_switches).
- **[Effective Go — Embedding](https://go.dev/doc/effective_go#embedding)** — the
  canonical explanation of why Go embeds instead of inheriting.
- **Jon Bodner, *Learning Go* (2nd ed., O'Reilly)** — the book being worked through.
  Chapter 7 covers types, methods, interfaces, and implicit interface satisfaction.

## Worked examples in the standard library

- **[`net/http.HandlerFunc`](https://pkg.go.dev/net/http#HandlerFunc)** — the
  canonical function adapter: a named func type with a method, so a plain function
  can satisfy an interface.
- **[`sort.Slice` / `sort.Interface`](https://pkg.go.dev/sort#Interface)** — adapting
  arbitrary data to an interface the algorithm demands.
- **[`io.WriteString`](https://pkg.go.dev/io#WriteString)** — the optional-behaviour probe in
  six lines: accept `io.Writer`, assert for `io.StringWriter`, fall back.
- **[Go blog: working with errors in Go 1.13](https://go.dev/blog/go1.13-errors)** — `%w`,
  `errors.Is`, `errors.As`. The clearest writing on the subject.

## Style

- **[Go Code Review Comments](https://go.dev/wiki/CodeReviewComments)** — the de-facto
  style guide; short, and worth rereading periodically.

## Communities

- **[r/golang](https://www.reddit.com/r/golang/)** — high-traffic, good for "is this
  idiomatic?" questions.
- **[Gophers Slack](https://invite.slack.golangbridge.org/)** — `#general` and
  `#newbies` are friendly and fast.
- **[golang-nuts](https://groups.google.com/g/golang-nuts)** — the mailing list;
  slower, but core team members answer language-semantics questions there.
