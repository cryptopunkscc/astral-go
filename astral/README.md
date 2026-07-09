# astral

The Go implementation of the astral object model: object types, codecs, and
the type registry. Network definitions and encodings live in the protocol
spec (see [Spec](#spec)); this package implements them.

## Object Model

- `Object` - the core interface: `ObjectType() string` plus
  `WriteTo`/`ReadFrom` for the payload (the type is outside the payload).
- `ObjectID` - `Size uint64` + 32-byte SHA-256 `Hash`; parses and renders the
  `data1` z-base32 text form.
- `Identity` - a secp256k1 public key; the zero value `Anyone` parses from
  `"anyone"` or the all-zero key.
- `Query` - `Nonce`, `Caller`, `Target`, `QueryString`.
- `Context` - wraps `context.Context` with identity, zone, and filters.
- `Blueprints` - the type registry; `New(typeName)` materializes a zero-value
  object of a registered type.

## Codecs

- `Objectify(&v)` builds the binary and JSON codecs for a struct through
  reflection, encoding fields in declaration order.
- Sized primitives carry explicit wire widths: `Int8`…`Int64`,
  `Uint8`…`Uint64`, `Float32`/`Float64`, `String8`…`String64`,
  `Bytes8`…`Bytes64`; plus `Bool`, `Time`, `Duration`, `Nonce`, `Nil`,
  `EOS`, `Ack`.
- `Adapt(v)` wraps a native Go value in its equivalent `Object`.

## Registration

- Every wire type registers its prototype with `astral.Add(&T{})` in its
  defining file.
- A blank import of `github.com/cryptopunkscc/astral-go` (package `pub`)
  registers the full wire surface: the astral primitives, `astral/log`, and
  all `api/` packages.

## Subpackages

- `channel` - typed object channels over byte streams; binary, JSON, text,
  and render formats. See [channel/README.md](channel/README.md).
- `fmt` - object formatting: `Format` renders format arguments as astral
  objects; `Printer` writes them through registered views.
- `log` - object-based logger; `styles`, `theme`, and `views` nested.

## Spec

- Definitions (object, object id, identity, zone, stamp):
  [core-definitions/](../.ai/system/core-definitions/).
- Encodings and framing: [topics/codec.md](../.ai/system/topics/codec.md).
