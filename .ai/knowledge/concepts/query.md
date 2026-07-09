# Query

A Query requests a bidirectional session with a named operation on a target
Identity. The `Query` wire object and its semantics — accept, reject codes,
the resulting Channel — are specified in
[query](../../system/core-definitions/query.md). This note covers the Go
routing types in `astral/` (`NewQuery` builds a `Query` with a random nonce).

## In-Flight Queries

* `InFlightQuery` wraps a `*Query` with `Extra sig.Map[string, any]`
  (`astral/in_flight_query.go`); `Launch(query)` constructs one.
* `Extra` carries routing metadata attached while the query is in flight.
* `OriginLocal` ("local") and `OriginNetwork` ("network") are values for
  `Extra["origin"]`.
* `IsNetwork()` is true only when the origin equals `OriginNetwork`.
* `IsLocal()` counts an unset or empty origin as local.

## Router

* `astral.Router` is the routing interface (`astral/router.go`):
  `RouteQuery(ctx *Context, q *InFlightQuery, w io.WriteCloser) (io.WriteCloser, error)`.
* `w` is the caller's response sink; the returned `WriteCloser` is the
  responder's request sink the caller writes into.
* A nil return with `ErrRouteNotFound` means the query was not handled;
  composite routers (`lib/routing` `PriorityRouter`) treat it as
  fall-through to the next router, stopping on success or `ErrRejected`.
* A `Router` may implement `HasRoutingPriority` to order itself among peers;
  lower values route first, absence implies `RoutingPriorityNormal`.
* The node-side routing pipeline — router registration, preprocessors,
  gateway relay — is astrald core; see astrald's query concept note.

## ErrRouteNotFound

* `ErrRouteNotFound` is an empty struct (`astral/err_route_not_found.go`).
* Its `Is` method matches any `*ErrRouteNotFound`, so match with
  `errors.Is(err, &astral.ErrRouteNotFound{})`, not by instance identity.
