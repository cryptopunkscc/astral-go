# Rules

## Boundary

- astral-go is the wire surface: primitives, wire types, op-name constants,
  protocol clients, client libraries.
- Node-side constructs belong to astrald: `Module` interfaces, engines,
  drivers, authorizer dispatch, `src/` implementations. Never add them here.
- astrald imports astral-go; astral-go never imports astrald.
- A wire-type change is a network-format decision: change the spec in
  `.ai/system` first, then this module, then consumers.
- Shared dependency versions stay pinned to astrald's (btcec, secp256k1).

## Context Discipline

- Implementation notes go in `.ai/knowledge/`; recipes in `.ai/patterns/`.
- Every note gets a keyword row in the index in `.ai/README.md`.
- Use the index before loading scoped files.
- Correct stale `.ai` context when found.

## Layout

- `astral/` - primitives, codecs, blueprints; `channel`, `fmt`, `log` nested.
- `sig/`, `streams/` - dependency-free utilities: signal-driven concurrency,
  stream helpers. They import nothing from this module.
- `api/<p>/` - one protocol: wire types + op constants; `client/`, where
  present, is its RPC client.
- `lib/` - app libraries: `apphost` (session), `apps` (serving),
  `astrald` (node client), `ipc`, `query`, `routing`.
- `pub.go` - the root registration aggregator; a blank import of
  `github.com/cryptopunkscc/astral-go` registers the full wire surface.
  Extend it when a new package registers blueprints.
- `examples/` - runnable app-developer programs, one concept each; every
  example builds with the module.

## Wire Types

- Every type defining `ObjectType() string` registers with `astral.Add(&T{})`
  in its defining file.
- Use `astral.Objectify` for `WriteTo`/`ReadFrom`. Prefer astral primitives
  for Objectify fields; platform-width `int`/`uint` are rejected — use sized
  types.
- Streaming ops end with `ch.Send(&astral.EOS{})`.
- Send stream errors with `ch.Send(astral.Err(err))`.
- Every `objects.Writer` must `Commit()` or `Discard()`.

## Project APIs

- Use `astral.Adapt(v)` to wrap a native Go value into an astral `Object`;
  do not hand-roll switch ladders. Pass-through for `Object`; `nil` and
  typed-nil pointers → `&Nil{}`; `error` → `NewError`. Default widths:
  `int`/`uint` → `Int64`/`Uint64`, `string` → `String32`. When the spec
  dictates a narrower width, Adapt is the wrong tool — dispatch on the
  spec first.
- Prefer `sig.Map`/`sig.Set`/`sig.Queue` over mutex + map/slice.
- Use `sig.RecvErr`/`sig.Recv`/`sig.Send` for context-aware channel ops.
- Logging: always `%v`. Levels: `Log` 0, `Logv(1)` verbose, `Logv(2)` debug.

## Code Shape

- Functions: one responsibility, max 50 lines, max 4 params.
- Packages: one concept. No `util`, `common`, `helpers`.
- Interfaces live at consumers. Prefer 1 method; 3+ is suspect.
- Naming: precise verbs, e.g. `delete`, `find`, `create`.

## Documentation Style

- Write `.ai` docs in the minimal English of `.ai/system/`.
- Declarative present tense. One fact per sentence or bullet.
- No motivation, hype, hedging, or meta-commentary.
- Backtick code identifiers. State defaults, limits, and terminators explicitly.
