# api/kcp

Wire surface of the `kcp` transport protocol: the `Endpoint` and
`EndpointLocalMapping` wire objects, op-name constants, and the typed
client. The KCP listener, dialing, and NAT port-mapping machinery are
astrald mod/kcp; op semantics are specified in
[protocols/kcp](../../system/protocols/kcp/README.md).

## Wire objects

- `Endpoint` (`endpoint.go`, `mod.kcp.endpoint`) is an `exonet.Endpoint`
  and `astral.Object`: `IP ip.IP`, `Port astral.Uint16`; `Network()`
  returns `"kcp"`.
- `Endpoint` text and JSON forms are specified in
  [mod.kcp.endpoint](../../system/protocols/kcp/types/mod.kcp.endpoint.md);
  `ParseEndpoint(s)` and `UnmarshalText` parse the text form and reject
  ports over 16 bits.
- `EndpointLocalMapping` (`endpoint_local_mapping.go`,
  `mod.kcp.endpoint_local_mapping`) maps a remote endpoint address to a
  local UDP port: `Address astral.String8`, `Port astral.Uint16`; encoding
  is specified in
  [mod.kcp.endpoint_local_mapping](../../system/protocols/kcp/types/mod.kcp.endpoint_local_mapping.md).
- Both objects encode via `astral.Objectify` and register their blueprints
  with `astral.Add` in package `init()`.
- `Endpoint.UDPAddr()` converts to `*net.UDPAddr`; `IsZero()` is safe on a
  nil receiver.

## Op names

`module.go`: `MethodNewEphemeralListener`, `MethodCloseEphemeralListener`,
`MethodSetEndpointLocalPort`, `MethodRemoveEndpointLocalPort`,
`MethodListEndpointLocalMappings` (`kcp.<op>`).

## Client

`client/` wraps a `lib/astrald` client with an optional target identity:
`New(targetID, client)` — a nil client falls back to `astrald.Default()`;
`WithTarget(target)` rebinds the target. There is no package-level default
client. One wrapper per op:

- `CreateEphemeralListener(ctx, port)`, `CloseEphemeralListener(ctx, port)`
  — expect `Ack`.
- `SetEndpointLocalPort(ctx, endpoint, localPort, replace)` — sends the
  endpoint's text address; `replace` controls whether an existing mapping
  is overwritten; expects `Ack`.
- `RemoveEndpointLocalPort(ctx, endpoint)` — expects `Ack`.
- `ListEndpointLocalMappings(ctx)` — collects `EndpointLocalMapping`
  objects until `EOS`.
