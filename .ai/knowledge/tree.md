# Tree

`tree.Value[T]` (`api/tree/value.go`) is a typed, observable cell bound to a
node of a tree store. The store itself (paths, mounts, the default DB-backed
node) is served by astrald mod/tree; the network ops are specified in
[protocols/tree](../system/protocols/tree/README.md).

## Binding

* `Bind(ctx, node)` follows the backing `tree.Node` until the context is
  canceled: it blocks for the initial value, then applies every update the
  node delivers.
* `BindPath(ctx, node, path, create)` walks `path` from `node` with
  `tree.Query` and binds to the final node; `create` makes missing segments.

## Reads

* `Get()` is a one-shot read of the currently held value (the cache the
  binding maintains).
* `Follow(ctx)` returns a channel that delivers the current value
  immediately, then every change; the channel closes when the context is
  canceled.

## Writes

* `Set(ctx, v)` writes to the backing node; a nil value is stored as
  `astral.Nil`. When the node rejects the write, `Set` falls back to a
  local-only update; setting `NoLocal` suppresses the fallback and returns
  the node's error instead.
* `Clear(ctx)` resets the cell to T's zero value, persisting `astral.Nil`;
  use it for pointer-typed T (a typed-nil T slips past `Set`'s nil guard).

## Semantics

* A stored `astral.Nil` arrives at readers as T's zero value.
* An update of a type other than T resets the cache to T's zero value and is
  not delivered to followers.
* Persistence belongs to the backing node, not the cell: the default astrald
  tree node stores values in the node's database, so bound values survive
  restarts there.
* `Value[T]` implements `astral.Object` by delegating to the held value, so
  structs with `Value` fields work with `astral.Objectify`; `WriteTo` on an
  empty cell returns an error.
* The `NoInit` field is currently unused; `Bind` does not consult it.
