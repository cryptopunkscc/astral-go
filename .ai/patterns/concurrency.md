# Concurrency Patterns

Recipes for context-aware channel operations and shared mutable state using
the `sig` package. `sig` is dependency-free; it imports nothing from this
module.

## Context-Aware Channel Helpers

Use `sig` helpers instead of hand-written `select` on `ctx.Done()`.

```go
err := sig.RecvErr(ctx, errc)    // receive from an error channel
v, err := sig.Recv[T](ctx, ch)   // receive one value
err = sig.Send(ctx, ch, value)   // send one value
```

- Each helper returns `ctx.Err()` when the context ends first.
- `sig.RecvOk` additionally reports the receive's `ok` flag.

Source: `sig/chan.go`

## sig Collections

Prefer `sig.Map`, `sig.Set`, and `sig.Queue` over raw mutex plus map/slice for
shared mutable state.

```go
m := sig.Map[string, *Thing]{}   // zero value is ready
m.Set("key", thing)              // inserts only when the key is absent
v, ok := m.Get("key")
all := m.Clone()

s := sig.Set[*Thing]{}
s.Add(thing)                     // variadic; duplicates return ErrDuplicateItem
s.Remove(thing)

q := &sig.Queue[*Event]{}
q = q.Push(event)                       // producer keeps the returned tail
for e := range sig.Subscribe(ctx, q) {  // consumer; ends on ctx or Close
    handle(e)
}
```

Rules:

- `sig.Map.Set` never overwrites; it returns the existing value and `false`.
  Use `Replace` to overwrite.
- `sig.Set` rejects duplicates and preserves insertion order (`Clone`, `Sort`).
- `sig.Queue` is unbounded; `Push` returns the new tail — store it back.
  `Close` marks EOF and ends subscriptions.

Source: `sig/map.go`, `sig/set.go`, `sig/queue.go`; usage: `api/tree/value.go`,
`lib/routing/op_router.go`

## Ring And Pool

- `sig.Ring[T]` is a fixed-capacity thread-safe ring buffer. Construct with
  `sig.NewRing(capacity)`; capacity must be positive. `Push` on a full ring
  overwrites and returns the oldest value; `Pop`/`Peek` return `false` when
  empty.
- `sig.Pool` is a named counting semaphore. `Add(item, count)` grows an item's
  capacity; `Lock(names...)` blocks until all named items are acquired
  atomically; pair every `Lock` with `Unlock`.

```go
r, _ := sig.NewRing[Event](16)
oldest, overwritten := r.Push(event)

p := sig.NewPool()
p.Add("slot", 4)
p.Lock("slot")
defer p.Unlock("slot")
```

Source: `sig/ring.go`, `sig/pool.go`
