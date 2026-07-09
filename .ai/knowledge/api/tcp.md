# api/tcp

Wire type and op-name constants of the `tcp` transport protocol; `client/`
is the protocol's RPC client. Op semantics and the endpoint encoding live in
the spec: [protocols/tcp](../../system/protocols/tcp/). Dialing, servers,
and listener lifecycle are node-side: astrald mod/tcp.

## Endpoint

- `Endpoint` (`endpoint.go`) — object type `mod.tcp.endpoint`; fields `IP`
  (`ip.IP`) and `Port` (`astral.Uint16`); registers with `astral.Add`.
- Implements `exonet.Endpoint`: `Network()` is always `"tcp"`, `Address()`
  and `String()` return `<ip>:<port>` (`net.JoinHostPort`), `Pack()` is the
  wire form.
- Text and JSON forms are specified in
  [types/mod.tcp.endpoint](../../system/protocols/tcp/types/mod.tcp.endpoint.md).
- `ParseEndpoint(s)` parses `host:port`; parsing rejects ports that do not
  fit in 16 bits.
- `IsZero` is safe on a nil pointer; a nil `IP` also counts as zero.

## Op names

`module.go`: `MethodNewEphemeralListener` (`tcp.new_ephemeral_listener`),
`MethodCloseEphemeralListener` (`tcp.close_ephemeral_listener`).

## Client

`client/` wraps a `lib/astrald` client with an optional target identity:
`New(targetID, client)` (a nil client falls back to `astrald.Default()`),
`WithTarget` re-targets.

- `CreateEphemeralListener(ctx, port)` — queries
  `tcp.new_ephemeral_listener` with a `port` arg.
- `CloseEphemeralListener(ctx, port)` — queries
  `tcp.close_ephemeral_listener` with a `port` arg.

Both expect `Ack` and return a received `astral.ErrorMessage` as the error.
