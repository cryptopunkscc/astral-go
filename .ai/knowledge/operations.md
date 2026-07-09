# Operations

An Op is a named service invoked by a Query method string. Protocol meaning —
op names, parameters, modes — is specified in
[op](../system/core-definitions/op.md) and
[op-modes](../system/topics/op-modes.md). This note covers the handler
framework in `lib/routing` and `lib/query`.

## Handlers

* `routing.NewOp(fn)` wraps a handler function (`lib/routing/op.go`).
* Valid signatures:
  * `func(*astral.Context, *routing.IncomingQuery) error`
  * `func(*astral.Context, *routing.IncomingQuery, Args) error` — `Args` is a
    struct, or pointer to a struct, with exported fields.
* Any other shape fails with `ErrInvalidSignature`.
* The handler resolves the query exactly once: `Accept` (returns a Channel),
  `AcceptRaw`, `Reject`, or `RejectWithCode` (`lib/routing/incoming_query.go`).
* `Op.RouteQuery` runs the handler in a goroutine with a detached context; the
  caller blocks until the handler resolves the query, the context ends, or a
  5-second deadline rejects it. A query left unresolved is rejected when the
  handler returns.

## Args Structs

* `query.Parse` splits the query string into a path and a `map[string]string`
  of params ([query-string](../system/core-definitions/query-string.md)).
* Params fill the args struct by field name; exported names convert
  PascalCase to snake_case (`lib/query/editor.go`).
* Values decode via `encoding.TextUnmarshaler` when implemented, else by
  native kind (strings, ints, uints, floats, bool; `[]byte` from base64).
* Unknown params are silently ignored (`Editor.SetMany`).

## Query Tags

The `query` struct tag holds semicolon-separated directives
(`lib/query/field_tag.go`):

| Directive | Effect |
|---|---|
| `key:<name>` | expose the field as `<name>` instead of the snake_cased field name |
| `required` | reject the query before the handler runs when the param is missing |
| `skip` | exclude the field from parsing and specs |

* Fields are optional by default; requiredness is opt-in via
  `query:"required"`.
* Any other directive — including the `optional` seen in older code — parses
  into `FieldTag.Other`, which nothing reads; it has no effect.
* A missing required param rejects with `field required: <name>` before the
  handler runs (`lib/routing/op.go`).

## Registration

* `OpRouter` routes by matching the query string's path — the part before
  `?` — against named Ops (`lib/routing/op_router.go`).
* `AddOp(name, op)` registers one Op; duplicate names error.
* `AddStructPrefix(s, prefix)` registers every exported method of `s` whose
  name starts with `prefix` and whose signature `NewOp` accepts; the prefix is
  stripped and the rest converts PascalCase to snake_case. astrald modules
  call it as `AddStructPrefix(mod, "Op")`.
* `Spec()` lists every registered Op with its argument specs.
