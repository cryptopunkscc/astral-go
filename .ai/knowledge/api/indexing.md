# api/indexing

Wire messages, op-name constants, and sentinels of the `indexing` protocol
— change delivery from named object repositories to registered indexers.
`client/` is the protocol's RPC client. Protocol semantics (ack matching,
+1 cursor advance, retry backoff) live in the spec:
[protocols/indexing](../../system/protocols/indexing/).

## Wire messages

`messages.go`; all three register with `astral.Add` and encode via
`astral.Objectify`.

| Message | Fields | Object type |
|---|---|---|
| `IndexMsg` | `Repo String8`, `Version Uint64`, `ObjectID` | `indexing.index` |
| `UnindexMsg` | `Repo String8`, `Version Uint64`, `ObjectID` | `indexing.unindex` |
| `ChangeAckMsg` | `Repo String8`, `Version Uint64` | `indexing.ack` |

## Op names

`module.go`: `MethodRegisterIndexer`, `MethodSubscribe`,
`MethodRemoveIndex` (`indexing.<op>`). The specced `indexing.enable_repo`
op has no constant and no client wrapper.

## Sentinels

`errors.go`: `ErrIndexNotFound`, `ErrRepositoryNotFound`, `ErrAckMismatch`
(plain Go errors) and `ErrIndexingTemporarilyFailed` (an `astral.Error`,
sent on the wire to signal a retryable failure).
`IsIndexingTemporarilyFailed` matches by error string because
`astral.Error` does not support `errors.Is` unwrapping.

## Client

`client/` wraps a `lib/astrald` client with an optional target identity
(`New(targetID, client)`, `Default()`); a nil target routes to the local
node. Every wrapper also exists as a package-level function on the default
client.

- `RegisterIndexer(ctx, name)` — returns the indexer's `astral.Nonce`
- `Subscribe(ctx, nonce)` — opens the change stream, returns a
  `*Subscription`; a background goroutine closes it when `ctx` ends
- `RemoveIndex(ctx, nonce)` — deregisters the indexer, expects `Ack`

## Subscription

`Subscription` (`client/subscribe.go`) holds at most one pending change:

- `Next()` blocks for the next `IndexMsg`/`UnindexMsg` and marks it
  pending; calling `Next` with an unacknowledged change is an error; a
  server-sent `astral.Error` returns as an error.
- `Ack()` sends a `ChangeAckMsg` echoing the pending change's
  `(Repo, Version)` and clears it.
- `Fail()` sends `ErrIndexingTemporarilyFailed` and clears the pending
  change; the server retries the same change with backoff.
- `Close()` is idempotent.

The changelog, per-indexer cursors, and repo sync are node-side: astrald
mod/indexing.
