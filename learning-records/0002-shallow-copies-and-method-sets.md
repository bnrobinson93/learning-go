# 0002 — Addressability is a call-site convenience, not an interface rule

- **Date:** 2026-09-02
- **Chapter:** 7 — Types, Methods, and Interfaces
- **Status:** learned
- **Follows:** [0001](0001-adapters-over-widening-types.md)

## Context

`League` mixes receiver kinds: `MatchResult` takes `*League`, `Ranking` takes `League`. The
struct-embedding adapter from lesson 1 (`type RankerStruct struct { League }`) therefore
embeds a copy, and the question arose of whether it should embed `*League` instead.

## What was learned

1. **Two separate rules, routinely conflated.**
   - Calling a method: Go auto-inserts `&`/`*` when the operand is *addressable*.
   - Satisfying an interface: only the method set counts, and an interface value has no
     address, so the convenience never applies.
2. Method set of `T` holds value-receiver methods; `*T` holds both. If any required method has
   a pointer receiver, only `*T` implements the interface.
3. Not addressable: composite literals, map elements, function results. Hence
   `cannot call pointer method AddPlayer on Team`.
4. Struct copies are **shallow**. Brad predicted the embedded copy would isolate mutations; it
   only half does. Verified:
   - `rs.MatchResult(…)` does `l.Wins[name]++` → visible in the original (map table shared).
   - `rs.Reset()` does `l.Wins = …` → not visible (field reassignment stays in the copy).
5. Same shape in `for _, t := range l.Teams` — `t` is a copy, so `t.AddPlayer(…)` is discarded.
   Index instead: `l.Teams[i].AddPlayer(…)`, because slice elements are addressable.

## The non-obvious bit

The failure mode is not a crash and not a compile error — it is *some* mutations landing and
others silently vanishing, depending on whether the code writes through a reference field or
reassigns it. That is why "just copy it, it's safer" is bad instinct in Go.

## Correction to an earlier prediction

Brad's answer to the trap — "it's a copy, so it will not update the actual league" — was right
about the struct and wrong about the map. Worth re-testing later; this is the kind of thing
that decays.

## Exercise outcome (same day)

Both fixes landed: indexed `AddPlayer` mutates the real team, `RankerStruct` embeds `*League`,
and `Ranking` was moved to a pointer receiver so `League` no longer mixes receiver kinds.

One thing survived the cleanup and is worth remembering, because it is the shallow-copy rule
wearing a disguise: `RankingStrings` kept a **value** receiver, and it is bound as a **method
value** (`RankerAdapter(league.RankingStrings)`). A value-receiver method value copies the
receiver at bind time, so the adapter holds a snapshot. It looks live only because `Wins` is a
shared map; reassigning `league.Wins` leaves the adapter stale. Verified:

```
f := league.RankingStrings
league.Wins["Broncos"] = 2         → f() sees it       (write through shared map)
league.Wins = map[...]{"Jets": 9}  → f() does NOT      (field reassignment)
```

**Generalisation:** a method value freezes as much of the receiver as the receiver kind
implies. Pointer receiver, nothing frozen; value receiver, a shallow copy frozen at bind time.

## Open threads

- `RankingStrings` should take `*League` — both for consistency and to kill the snapshot.
- `errcheck`: the first `RankPrinter` call site still discards its error (the second was fixed).
- Not yet covered in chapter 7: type assertions and type switches, embedding vs inheritance,
  implicit interface satisfaction as a design tool ("accept interfaces, return structs").
