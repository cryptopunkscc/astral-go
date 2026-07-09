# api/nat

Wire surface of the `nat` protocol — UDP hole punching between nodes: the
`Hole` and `Endpoint` wire objects, the punch and consume-hole protocol
signals, the puncher abstraction, op-name constants, and the typed client.
The cone-NAT probe socket, hole pool, keepalive, and enablement logic are
astrald mod/nat.

## Wire objects

- `Hole` (`hole.go`, `nat.hole`) is a connected endpoint pair from a
  successful punch: `Nonce`, `ActiveIdentity`/`ActiveEndpoint`,
  `PassiveIdentity`/`PassiveEndpoint`, `CreatedAt`. Helpers resolve the
  role-dependent view: `RemoteIdentity(self)`, `GetLocalAddr(self)`,
  `GetRemoteAddr(self)`, `MatchesPeer(peer)`.
- `Endpoint` (`endpoint.go`, `nat.endpoint`) is a UDP address: `IP ip.IP`,
  `Port astral.Uint16`; `Network()` returns `"utp"`; `MarshalText` and
  `String` return `host:port`; `ParseEndpoint` rejects ports over 16 bits;
  `UDPAddr()` converts to `*net.UDPAddr`.
- Package `init()` registers the blueprints of all four wire objects
  (`Hole`, `Endpoint`, `PunchSignal`, `ConsumeHoleSignal`) with
  `astral.Add`.

## Protocol signals

- `PunchSignal` (`punch_signal.go`, `nat.punch_signal`): `Signal`,
  `Session`, `IP`, `Port`, `PairNonce`. Signal types: `offer`, `answer`,
  `ready`, `go`, `result` — the initiator sends offer, the passive side
  answers and signals ready, the initiator signals go (both sides punch
  simultaneously), the initiator sends result.
- `ConsumeHoleSignal` (`consume_hole_signal.go`, `nat.consume_hole_signal`):
  `Signal`, `Pair`, `Ok`, `Error`. Signal types: `lock`, `locked`, `take`,
  `taken` — the two-phase handshake that claims a hole.
- `ExpectConsumeHoleSignal(pair, type, on)` builds a channel handler that
  validates pair and signal type; `HandleFailedConsumeHoleSignal` converts
  `Ok=false` into the signal's error.

## Punch types

- `Puncher` (`puncher.go`) abstracts the UDP punch: `Open()` binds a local
  socket and returns its port, `HolePunch(ctx, peerIP, peerPort)` returns a
  `PunchResult`, `Session()` is the opaque byte sequence stamped into punch
  packets, `Close()` releases the socket.
- `PunchResult` reports only what this side observed of the peer:
  `LocalPort`, `RemoteIP`, `RemotePort`.
- `PunchProtocol` (`punch_protocol.go`) is the pure signalling state
  machine: builds `OfferSignal`/`AnswerSignal`/`ReadySignal`/`GoSignal`/
  `ResultSignal`, records peer state via `OnOffer`/`OnAnswer`/`OnResult`,
  and `ExpectSignal` rejects any non-offer signal whose session differs
  from the established one. `SetPunchResult` marks the local side active
  and records the observed remote endpoint as the passive endpoint.

## Op names

`module.go`: `MethodPunch`, `MethodNodePunch`, `MethodNodeConsumeHole`,
`MethodListHoles`, `MethodSetEnabled` (`nat.<op>`).

## Client

`client/` wraps a `lib/astrald` client with an optional target identity
(`New(targetID, client)`, `Default()`, `SetDefault`); a nil target routes
to the local node.

- `Punch(ctx, target)` (`traverse.go`) — opens `nat.punch`, returns the
  resulting `*nat.Hole`.
- `NodePunch(ctx, target, localIP, puncher)` — drives the active half of
  the offer/answer/ready/go/result exchange over a `nat.node_punch`
  channel, delegates the UDP burst to the supplied `Puncher`, generates the
  hole `Nonce`, and returns the completed `Hole`.
- `NodeConsumeHole(ctx, pair, target)` — with a target, acts as initiator
  and waits for `Ack`; with a nil target, acts as responder and drives
  lock → locked → take → taken.
- `ListHoles(ctx, with)` — optional peer filter; collects `*nat.Hole` until
  `EOS`.
- `SetEnabled(ctx, enabled)` — expects `Ack`.
