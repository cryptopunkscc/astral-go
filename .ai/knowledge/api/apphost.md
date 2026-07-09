# api/apphost

Wire types and op-name constants of the `apphost` protocol — the on-device
guest protocol that lets local apps route queries through the node. `client/`
is the protocol's RPC client. Framing, handshake, and message-flow semantics
live in the spec: [astral-ipc](../../system/topics/astral-ipc.md),
[ws-transport](../../system/topics/ws-transport.md),
[protocols/apphost](../../system/protocols/apphost/).

## Guest wire messages

One file per message; every type registers with `astral.Add` and encodes as
`mod.apphost.<name>`.

| Message | Meaning |
|---|---|
| `HostInfoMsg` | host greeting: node `Identity` + `Alias` |
| `AuthTokenMsg` | guest presents an access token (`String8`) |
| `AuthSuccessMsg` | auth accepted; carries the resolved `GuestID` |
| `ErrorMsg` | coded error; `Code String8`, implements `error` |
| `RouteQueryMsg` | outbound query: `Nonce`, `Caller`, `Target`, `Query`, `Zone`, `Filters` |
| `QueryAcceptedMsg` | query accepted; raw byte stream follows |
| `QueryRejectedMsg` | query rejected with a `Uint8` code |
| `RegisterHandlerMsg` | register an IPC query handler: `Identity`, `Endpoint`, `AuthToken` nonce |
| `BindMsg` | bind a handler token (`Nonce`) to the guest session |
| `HandleQueryMsg` | host-to-handler query delivery over a fresh IPC conn |
| `RegisterServiceMsg` | WS-only: register an inbound-query service for `Identity` |
| `IncomingQueryMsg` | WS-only: notify the service of a pending inbound query |
| `AttachQueryMsg` | WS-only: attach a new WS conn to a pending query by `QueryID` |
| `RejectIncomingMsg` | WS-only: reject a pending inbound query with a `Uint8` code |
| `PingMsg` | keepalive; zero-byte payload |

`ErrorMsg` codes are the `ErrCode*` constants in `error_msg.go`: `auth_failed`,
`denied`, `route_not_found`, `internal_error`, `protocol_error`, `timeout`,
`canceled`, `target_not_allowed`.

## Objects

- `AccessToken` (`access_token.go`) — `Identity`, `Token String8`,
  `ExpiresAt Time`; object type `apphost.access_token` (no `mod.` prefix,
  unlike the messages).
- `App` (`app.go`) — installed-app record: `AppID`, `HostID`, `InstalledAt`.

## Op names

`module.go` defines `MethodCreateToken`, `MethodListTokens`,
`MethodRegisterHandler`, `MethodCancel`, `MethodBind`, `MethodNewAppContract`,
`MethodSignAppContract`, `MethodInstallApp`, `MethodHoldObject`,
`MethodUnholdObject`, `MethodListHeldObjects` (`apphost.<op>`), plus the
`ErrProtocolError` sentinel. The specced `apphost.register` and
`apphost.whoami` ops have no constants.

## Client

`client/` wraps a `lib/astrald` client with an optional target identity
(`New(targetID, client)`, `Default()`); every wrapper also exists as a
package-level function on the default client.

- `CreateToken`, `ListTokens` — access-token management
- `RegisterHandler` — register an IPC handler endpoint, expect `Ack`
- `Bind` — open the `apphost.bind` channel, wait for `Ack`, return the open
  channel; the caller closes it
- `NewAppContract`, `SignAppContract` — app relay contracts (`auth.Contract`)
- `HoldObject`, `UnholdObject` — object holds
- `ListHeldObjects` — streams `ObjectID`s on a Go channel; the returned
  `*error` is populated only after the stream closes

No wrappers exist for `cancel`, `install_app`, `register`, or `whoami`.

The node side (guest serving, HTTP/WS servers, token and hold persistence,
routing) is astrald mod/apphost.
