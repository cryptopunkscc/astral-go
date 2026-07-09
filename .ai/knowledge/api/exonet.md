# api/exonet

The `Endpoint` interface — a dialable address on a network. The package
holds only this interface (`module.go`); the transport registry (dialers,
parsers, unpackers) and the raw `Conn` contract are astrald mod/exonet.
The exonet concept — the untrusted carrier networks nodes exchange data
over — is defined in
[core-definitions/exonet](../../system/core-definitions/exonet.md).

## Endpoint

- An `Endpoint` is an `astral.Object` exposing `Network()` (transport
  name), `Address()` (text form of the address), and `Pack()` (binary
  form).
- Endpoints are serializable objects; published endpoints travel wrapped
  in `nodes.EndpointWithTTL`, which pairs an `Endpoint` with an optional
  TTL and itself implements `Endpoint` by delegation — see
  [mod.nodes.endpoint_with_ttl](../../system/protocols/nodes/types/mod.nodes.endpoint_with_ttl.md).

## Implementations

- In this module: the endpoint types of `api/tcp`, `api/kcp`, `api/utp`,
  `api/tor`, `api/gateway`, and `api/nat`, plus
  `api/nodes.EndpointWithTTL`.
