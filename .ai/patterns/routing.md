# Routing Patterns

Use when implementing `astral.Router`. The Router interface and in-flight
queries are described in [query](../knowledge/query.md); node-side
zone gating is astrald's routing pattern.

## RouteQuery Return Values

`astral.Router.RouteQuery(ctx *astral.Context, q *astral.InFlightQuery, w io.WriteCloser)`
returns `(io.WriteCloser, error)`. `w` is the caller's response sink; the
returned `WriteCloser` is the responder's request sink. Return the first
matching result. Never fall through with `nil, nil`.

```go
func (r *MyRouter) RouteQuery(ctx *astral.Context, q *astral.InFlightQuery, w io.WriteCloser) (io.WriteCloser, error) {
    if !r.matches(q) {
        return query.RouteNotFound()
    }
    if !r.authorized(q) {
        return query.Reject()
    }
    return query.Accept(q, w, func(conn astral.Conn) {
        defer conn.Close()
    })
}
```

| Situation | Return |
|---|---|
| Not our query | `query.RouteNotFound()` |
| Explicit refusal | `query.Reject()` (or `query.RejectWithCode(code)`) |
| Accepted | `query.Accept(q, w, handler)` |
| Never | `nil, nil` |

* `query.RouteNotFound` and `query.Reject` take no arguments. `Reject()`
  rejects with `astral.DefaultRejectCode` (1); `RejectWithCode(code)` panics
  on code 0.
* `query.Accept(q, w, handler)` runs `handler` in a new goroutine with the
  resulting `astral.Conn`.
* `ErrRouteNotFound` is non-terminal: composite routers (`lib/routing`
  `PriorityRouter`) fall through to the next router, stopping only on
  success or `ErrRejected`.
* `ErrRouteNotFound` carries no router reference; routers identify
  themselves through their own `String()` or `Name` (e.g. `PriorityRouter`).

Source: `astral/router.go`, `lib/query/route.go`, `astral/err_route_not_found.go`
