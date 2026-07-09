# api/services

Wire surface of service discovery: the `Update` wire object, op-name
constants, and the discovery client. Aggregation, discoverer registration,
and the sync cache are astrald mod/services; op semantics are specified in
[protocols/services](../../system/protocols/services/README.md).

## Wire surface

- `module.go` declares `MethodDiscover` (`services.discover`) and
  `MethodSync` (`services.sync`).
- `Update` (`update.go`) is the `services.update` wire object: `Available`,
  `Name`, `ProviderID`, `Info`; it registers with `astral.Add` and encodes
  via `astral.Objectify`. Encoding is specified in
  [services.update](../../system/protocols/services/types/services.update.md).
- `services.sync` has no client wrapper; only the constant is defined here.

## Discovery client

`client.Discover(ctx, follow)` (`api/services/client/services.go`) opens
`services.discover` on the target and returns a `<-chan *services.Update`.

- Snapshot phase: forwards `Update` objects until the server's `EOS`.
- `follow=false`: the channel closes after the snapshot.
- `follow=true`: one nil sentinel marks the snapshot/live boundary, then live
  updates are forwarded until the context is canceled or the server closes.
- Receiving an object of any other type ends the stream.
- `New(targetID, astralClient)` targets any node; a nil client falls back to
  `astrald.Default()`; the package-level `Discover` uses `Default()`.
