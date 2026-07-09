# Zone

`Zone` is a `uint8` bitmask scoping where a query may be resolved
(`astral/zone.go`); its object type is `zone`. Zone meanings and the
network-strip rule for untrusted queries are specified in
[Zone](../../system/core-definitions/zone.md).

## Bits

| Zone          | Bit | Letter |
|---------------|-----|--------|
| `ZoneDevice`  | `1` | `d`    |
| `ZoneVirtual` | `2` | `v`    |
| `ZoneNetwork` | `4` | `n`    |

* `ZoneDefault = ZoneAll = ZoneDevice|ZoneVirtual|ZoneNetwork`.
* `zone.Is(check)` tests that all bits in `check` are set.

## Text Form

* `String()` concatenates the letters of the set bits in `d`, `v`, `n` order.
* `Zones(s)` parses the letters back; unknown characters are ignored.
* JSON marshals and unmarshals the text form as a string.

## Context

`astral.Context` carries the zone (`astral/context.go`):

* `NewContext` returns a context with `ZoneDefault`.
* `ctx.Zone()` reads the current zone.
* `WithZone(z)` replaces, `IncludeZone(z)` adds, `ExcludeZone(z)` removes,
  `LimitZone(z)` intersects; each returns a clone.
