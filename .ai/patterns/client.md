# Operation Patterns

Recipes for writing protocol clients under `api/<p>/client/`. The handler
framework — signatures, args structs, query tags, registration — is covered
in [operations](../knowledge/operations.md); the `api/` tree
conventions in [api](../knowledge/api.md).

## Module Client

A protocol's client package lives at `api/<p>/client/`, one exported method
per op, mirroring the node module's `src/op_*.go` handlers in astrald.

- The package declares the protocol's name (`package nat` in
  `api/nat/client/`), so importers alias it, e.g. `natclient`.
- Methods invoke the `Method*` op-name constants from the api package root
  (e.g. `objects.MethodStore`).

```go
type Client struct {
    astral   *astrald.Client
    targetID *astral.Identity
}

var defaultClient *Client

func New(targetID *astral.Identity, a *astrald.Client) *Client {
    if a == nil {
        a = astrald.Default()
    }
    return &Client{astral: a, targetID: targetID}
}

func Default() *Client {
    if defaultClient == nil {
        defaultClient = New(nil, nil)
    }
    return defaultClient
}

func SetDefault(client *Client) {
    defaultClient = client
}

func (client *Client) queryCh(ctx *astral.Context, method string, args any, cfg ...channel.ConfigFunc) (*channel.Channel, error) {
    return client.astral.WithTarget(client.targetID).QueryChannel(ctx, method, args, cfg...)
}
```

- A nil `targetID` routes to the local node; a nil `*astrald.Client` falls
  back to `astrald.Default()` (routed through apphost).
- `Default()` builds the package default lazily; `SetDefault` replaces it.
- Every op goes through `queryCh`:
  `WithTarget(targetID).QueryChannel(ctx, method, args, cfg...)`.

Source: `api/nat/client/client.go` (same shape in `api/objects/client/client.go`,
`api/auth/client/client.go`)

## Client Operation File

Keep one client operation per file, named after the op (`store.go`,
`sign_contract.go`). Each op gets a `*Client` method plus a package-level
function delegating to `Default()`.

```go
func (c *Client) DoThing(ctx *astral.Context, contract *foo.Contract) (result *foo.SignedContract, err error) {
    ch, err := c.queryCh(ctx, foo.MethodDoThing, nil)
    if err != nil {
        return
    }
    defer ch.Close()

    if err = ch.Send(contract); err != nil {
        return
    }

    err = ch.Switch(channel.Expect(&result), channel.PassErrors)
    return
}

func DoThing(ctx *astral.Context, contract *foo.Contract) (*foo.SignedContract, error) {
    return Default().DoThing(ctx, contract)
}
```

- Pass scalar parameters as query args (`query.Args{"repo": repo}`); send
  object payloads with `ch.Send` after the channel is established.
- Read the result with `ch.Switch(channel.Expect(&result), channel.PassErrors)`;
  `PassErrors` surfaces received error objects as Go errors.

Source: `api/objects/client/store.go`, `api/auth/client/sign_contract.go`
