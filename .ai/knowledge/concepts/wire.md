# Wire

astral-go realizes the object codec through reflection. The wire framings —
binary, canonical, JSON, text — are specified by the spec
([codec](../../system/topics/codec.md),
[json-encoding](../../system/topics/json-encoding.md),
[text-encoding](../../system/topics/text-encoding.md)); a stable `ObjectID`
comes only from the canonical form. This note covers how the Go reflector
turns a struct into those framings.

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
