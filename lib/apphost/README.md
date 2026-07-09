# apphost

Client library for the node's apphost service. Protocol spec:
[astral-docs/protocols/apphost](https://github.com/cryptopunkscc/astral-docs/tree/master/protocols/apphost).

## Basic usage

```go
ctx := astral.NewContext(nil)

host, err := apphost.Connect(ctx, apphost.DefaultEndpoint)
if err != nil {
	return err
}

fmt.Printf("connected to host %v (%v)\n", host.HostAlias(), host.HostID())

if err := host.AuthToken(token); err != nil {
	return err
}

fmt.Printf("authenticated as %v\n", host.GuestID())

conn, err := host.RouteQuery(
	astral.Launch(query.New(nil, nil, "user.info", nil)),
	ctx.Zone(),
	nil,
)
```
