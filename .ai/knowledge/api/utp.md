# api/utp

The `utp.Endpoint` wire object — the address type of the uTP (Micro
Transport Protocol over UDP) transport. The package holds only
`endpoint.go`: no op-name constants, no client. The uTP listener and
dialing are astrald mod/utp.

## Endpoint

- `Endpoint` (`mod.utp.endpoint`) is an `exonet.Endpoint` and
  `astral.Object`: `IP ip.IP`, `Port astral.Uint16`.
- `Network()` is always `"utp"`.
- Text form: `MarshalText` and `String` both return `host:port` without
  the network prefix; `UnmarshalText` parses that form and rejects ports
  over 16 bits.
- JSON form: `MarshalJSON` encodes the same `host:port` string;
  `UnmarshalJSON` decodes it via `ParseEndpoint`.
- Wire form: `WriteTo`/`ReadFrom` via `astral.Objectify`; `Pack()` returns
  the binary form, nil on encoding error.
- `ParseEndpoint(s)` parses `host:port`; `UDPAddr()` converts to
  `*net.UDPAddr`; `IsZero()` is safe on a nil receiver.
- Package `init()` registers the blueprint with `astral.Add`.
