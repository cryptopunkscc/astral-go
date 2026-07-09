# api/tree

Wire surface of the tree object store: the `tree.Node` contract, path
traversal, the shared op handlers, and the remote client. The default store
(DB nodes, mounts) is served by astrald mod/tree; op semantics are specified
in [protocols/tree](../../system/protocols/tree/README.md).

## Node and traversal

- `tree.Node` (`api/tree/node.go`) is the store contract: `Get(ctx, follow)`,
  `Set`, `Delete`, `Sub`, `Create`. With `follow=true`, `Get`'s channel keeps
  delivering updates until the context is canceled.
- `tree.Query(ctx, root, path, create)` walks path segments through `Sub`;
  empty segments are skipped. With `create=true` missing segments are made
  with `Create`; `ErrAlreadyExists` is tolerated as a concurrent create and
  the walk continues.
- Typed helpers: `tree.Get[T]` is a one-shot read with a typecast;
  `tree.Follow[T]` streams values of type T and ends the stream with
  `ErrTypeMismatch` (via the returned error pointer) on any other type.
- `module.go` declares the op-name constants `MethodGet`, `MethodSet`,
  `MethodDelete`, `MethodList`, `MethodMountRemote`, `MethodUnmount`.
- `errors.go` sentinels: `ErrNodeHasSubnodes`, `ErrUnsupported` (wire
  errors), `ErrTypeMismatch`, `ErrAlreadyExists`; `ErrNoValue`
  (`err_no_value.go`) is both an error and a wire object.
- `tree.Value[T]` (`value.go`) binds a typed, observable cell to a node; see
  [concepts/tree](../concepts/tree.md).

## Op handlers

`NodeOps` (`api/tree/client/server.go`) serves the tree ops over any
`tree.Node` through `lib/routing`; astrald mod/tree delegates its `tree.get`
… `tree.list` handlers to it. All handlers accept `in`/`out` channel-format
args. Mount/unmount have no `NodeOps` counterpart; they are node-side only.

- `Get`: traverses without create; `follow=true` streams value updates; a
  background `Receive` on the accepted channel cancels the stream when the
  caller closes its side (EOF); ends with `EOS`.
- `Set`, single mode (`value` non-empty): `setSingle` instantiates
  `astral.New(type)` and parses `value` as text via
  `encoding.TextUnmarshaler`; an empty `type` is inferred from the node's
  current value and fails when the node holds none; the node is auto-created
  only when `type` is explicit; answers `Ack`.
- `Set`, batch mode (`value` empty): `setBatch` traverses with create, then
  reads objects until `EOS`, answering each `node.Set` with `Ack` or a wire
  error.
- `Delete`: `recursive=true` walks `Sub` depth-first and deletes leaves
  first, recursing through remote mounts; otherwise a single `node.Delete`;
  answers `Ack`.
- `List`: an empty path defaults to `/`; streams child names alpha-sorted as
  `String8`; ends with `EOS`.

## Remote client

`Client` and `Node` (`api/tree/client/client.go`, `node.go`):

- `New(targetID, astralClient)` targets any node; a nil client falls back to
  `astrald.Default()`. `Default()`/`SetDefault` manage a package default;
  `Root()` returns the target's root node.
- `Node` implements `tree.Node` by translating operations into query-channel
  calls against the target: `Get` → `tree.get` (with `follow`), `Set` →
  `tree.set` expecting `Ack`, `Delete` → `tree.delete` expecting `Ack`,
  `Sub` → `tree.list` collecting `String8` names.
- `Get(follow=true)` receives the initial value, then forwards updates until
  `EOS`, the context ends, or the server closes.
- `Create` issues `tree.set` without sending a value; the empty set creates
  the node server-side.
- Because `Node` is a `tree.Node`, it composes with `Query`, `NodeOps`, and
  `tree.Value`.
