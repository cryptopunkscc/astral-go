# Wire

## Encodings

| Format    | Framing                                     | Used for                  |
|-----------|---------------------------------------------|---------------------------|
| binary    | `String8(type)` + `Bytes32(payload)`        | default channel transport |
| json      | `{"Type":"...","Object":{...}}\n`           | debugging, CLI tooling    |
| text      | `#[type] value\n` or `#[type]:base64\n`     | human-readable output     |
| canonical | `Stamp(4b)` + `String8(type)` + raw payload | storage, ObjectID hashing |

Invariant: only canonical encoding produces a stable `ObjectID`. Do not use
binary framing for storage.

The spec defines the framings: [codec](../../system/topics/codec.md) (binary
and canonical), [json-encoding](../../system/topics/json-encoding.md),
[text-encoding](../../system/topics/text-encoding.md).

## Objectify Fields

`astral.Objectify(&v)` (`astral/objectify.go`) reflects a non-nil pointer into
binary, JSON, and a derived `ObjectType()`. The reflector reads and writes
struct fields in declaration order (binary encoding is positional — see
[structure](../../system/core-definitions/structure.md)).

Supported kinds:

* numeric: `Int8`...`Int64`, `Uint8`...`Uint64`, `Float32/64` — sized only;
  platform-width `int`/`uint` are rejected
* `Bool`, `String`, `Ptr`, `Slice`, `Array`, `Map`, nested `Struct`,
  `Interface`
* any field type that implements the codec itself, such as the astral
  primitives `Duration`, `Nonce`, `*Identity`, `ObjectID`, `Zone`

Interface fields typed as `astral.Object` use dynamic framing (`String8` type
name + payload; a zero-length type means nil). Types decoded into them must be
registered with `astral.Add`, kept in a `Blueprints` registry
(`astral/blueprints.go`).

## Helpers

* `astral.Stringify(v)` (`astral/stringify.go`) returns `Stringer.String()`,
  falls back to `TextMarshaler.MarshalText`, plain `string`, then `%v`.
* `astral.New(typeName)` (`astral/blueprints.go`) returns a zero-value object
  from the default `Blueprints` or nil.
