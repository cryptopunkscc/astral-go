# astral/log

Structured logging: a log entry is a wire object, so log output can cross the
network like any other object.

## Surface

- `Logger` writes entries. `Log` logs at level 0; `Logv(level, ...)` at a
  higher (more verbose) level. `Info*`/`Error*` are aliases and carry no
  severity.
- Format args pass through `astral.Adapt`, so entries carry typed objects,
  not strings; values Adapt cannot map render as `String32` (via a `fmt.View`
  or `astral.Stringify`).
- `SetPrefix(obj...)` and `Tag(tag)` return child loggers that prepend
  objects to every entry; children share the root's filter and `EntryLogger`
  set.
- `SetFilter` gates emission at the root; `AddLogger`/`RemoveLogger` attach
  `EntryLogger` sinks that receive every entry.
- `Entry` (`astrald.log.entry`: Origin, Level, Time, Objects) and `Tag`
  (`astrald.log.tag`, a `String8`) register with `astral.Add`.
- `ToSnakeCase` converts PascalCase to snake_case; `lib/routing`'s `OpRouter`
  uses it to derive op names.

## styles/, theme/, views/

- `styles/` — reusable color, gradient, and style helpers for terminal
  rendering: HSL `Color` (named palette; `Bri`/`Sat`/`Hue`/`Tetrad`
  transforms), `Gradient`, lipgloss-backed `Style`, `ColorFromString`
  (deterministic per-string color), `StringView` (styled `String32`).
- `theme/` — the shared palette built on `styles`: `Primary` plus its tetrad,
  and per-kind colors (`Error`, `Identity`, `ObjectID`, `Op`, `Type`, ...).
  astrald mod/log views and other node views import this palette.
- `views/` — base `astral/fmt` views registered in `init()`: `EntryView`,
  `IdentityView` (fingerprint), `QueryView`, `ShortTimeView`. Richer
  renderers live in astrald mod/log.
