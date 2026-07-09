# api/tor

Wire objects of the `tor` network: the onion endpoint and its digest.
astral-docs has no tor protocol spec yet, so the encodings below are defined
by this package. The SOCKS5 dialer, hidden-service lifecycle, and key
persistence are node-side: astrald mod/tor.

## Digest

- `Digest` (`digest.go`) — a `[]byte` object, type `mod.tor.digest`;
  registers with `astral.Add`.
- Wire form is exactly `DigestSize` (35) raw bytes, no length prefix;
  `ReadFrom` reads exactly 35 bytes and fails on a short read.
- Text form is lowercase base32 with a `.onion` suffix (`MarshalText`,
  `String`); `UnmarshalText` is case-insensitive, accepts an optional
  `.onion` suffix, and rejects decoded lengths other than 35.
- JSON encodes the text form. `DigestFromString` parses either variant.

## Endpoint

- `Endpoint` (`endpoint.go`) — fields `Digest` and `Port`
  (`astral.Uint16`); object type `mod.tor.endpoint`; registers with
  `astral.Add`.
- Implements the `exonet.Endpoint` methods: `Network()` returns
  `ModuleName`, `Address()` returns `<digest>.onion:<port>`, `Pack()` is the
  wire form.
- Text form is `<digest>.onion:<port>` (`MarshalText`); JSON encodes
  `Address()`; `String()` prefixes the address with the network name
  (`tor:<address>`).
- The zero endpoint's `Address()` is `"unknown"`; `UnmarshalText` maps
  `"unknown"` back to the zero value, so the JSON form round-trips the zero
  value (`MarshalText` does not — it yields `.onion:0`).
- `IsZero` is safe on a nil pointer; an empty digest counts as zero.

## Names

- `ModuleName` = `"tor"` (`module.go`) — the exonet network name.
- The package declares no `Method*` op constants and has no `client/`
  subpackage.
