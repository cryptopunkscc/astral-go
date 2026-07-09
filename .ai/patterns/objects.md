# Object Patterns

Recipes for defining typed payloads (wire objects).

## Object Definition

Define a typed payload with:

* An `ObjectType` method.
* `WriteTo` and `ReadFrom` methods backed by `astral.Objectify`.
* Registration via `astral.Add` in `init`.

```go
var _ astral.Object = &Note{}

type Note struct {
    Author astral.String8
    Text   astral.String16
}

func (n Note) ObjectType() string { return "example.note" }

func (n Note) WriteTo(w io.Writer) (int64, error) {
    return astral.Objectify(&n).WriteTo(w)
}

func (n *Note) ReadFrom(r io.Reader) (int64, error) {
    return astral.Objectify(n).ReadFrom(r)
}

func init() { _ = astral.Add(&Note{}) }
```

Rules:

* `astral.Objectify` requires a non-nil pointer; it panics otherwise.
* Prefer astral primitives for fields; platform-width `int`/`uint` are
  rejected — use sized types.
* `astral.Add` registers the prototype with the default `Blueprints`, making
  the type constructible by name via `astral.New` and decodable into
  `astral.Object` interface fields.
* When the type lives in a new package, extend `pub.go` so a blank import of
  the module root registers it.

Source: `examples/custom-wire-type/main.go`, `api/apphost/bind_msg.go`,
`astral/objectify.go`, `astral/blueprints.go`

Add JSON support only when needed:

```go
func (n Note) MarshalJSON() ([]byte, error)  { return astral.Objectify(&n).MarshalJSON() }
func (n *Note) UnmarshalJSON(b []byte) error { return astral.Objectify(n).UnmarshalJSON(b) }
```

Source: `api/objects/probe.go`

Field type reference: `.ai/knowledge/concepts/wire.md`.
