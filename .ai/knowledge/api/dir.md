# api/dir

Wire objects and op-name constants of the `dir` protocol — the identity
directory: alias-to-identity mapping, name resolution, and named target
filters. `client/` is the protocol's RPC client. Op semantics and encodings
live in the spec: [protocols/dir](../../system/protocols/dir/).

## Wire objects

- `Alias` (`alias.go`) — human-readable node name; a `String8` on the wire,
  object type `mod.dir.alias`.
- `AliasMap` (`alias_map.go`) — alias-to-`Identity` map, object type
  `mod.dir.alias_map`; built in one shot by the producer, read-only after
  return.

## Op names

`module.go`: `MethodAliasMap`, `MethodApplyFilters`, `MethodFilters`,
`MethodGetAlias`, `MethodResolve`, `MethodSetAlias` (`dir.<op>`).

## Client

`client/` wraps a `lib/astrald` client with an optional target identity
(`New(targetID, client)`, `Default()`); every wrapper also exists as a
package-level function on the default client.

- `GetAlias`, `SetAlias` — alias lookup and mutation
- `AliasMap` — full alias-to-identity map
- `ResolveIdentity` — name to `Identity`; parses a literal public key
  client-side before querying `dir.resolve`
- `Filters` — registered filter names
- `ApplyFilters` — tests an identity against named server-side filters

`Client.EnableCache` memoizes resolved identities and aliases for the
client's lifetime; `ClearCache` resets both caches. Caching is off by
default.

## Invariants

- Known bug: `Client.ApplyFilters` queries `dir.MethodSetAlias` instead of
  `dir.MethodApplyFilters` (`client/apply_filters.go:19`); the
  `dir.apply_filters` op is broken via the client until the code is fixed.
- `ResolveIdentity` returns a parsed literal identity without querying the
  node.

The alias table, resolver chain, and filter gate are node-side: astrald
mod/dir.
