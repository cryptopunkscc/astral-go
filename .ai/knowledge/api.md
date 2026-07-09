# api

Conventions of the `api/` tree — the per-protocol wire surface — and the
`pub.go` registration aggregator. Per-protocol details live in
`knowledge/api/<p>.md`; node-side `Module` interfaces and op implementations
are astrald mod/<p>.

## Package shape

- One package per protocol; 20 packages.
- Wire types and `Method*` op-name constants sit at the package root; the
  constants live in `module.go` (e.g. `api/objects/module.go`,
  `api/tcp/module.go`).
- Not every package has a `module.go`: `ip` and `utp` have none — their wire
  types sit in other root files (`ip/ip.go`, `utp/endpoint.go`).
- Wire types register with `astral.Add(&T{})` in their defining files
  (`.ai/rules.md` "Wire Types").
- Interfaces consumed by `lib/apps` hooks are defined at the package root:
  `objects.Finder`/`Searcher`/`Describer` back `apps.WithObjectFinder`/
  `WithObjectSearcher`/`WithObjectDescriber`; `exonet.Endpoint` is the
  dialable-address contract the other endpoint types implement.
- `client/` is the protocol's RPC client, present in 16 of 20 packages;
  absent in `exonet`, `ip`, `tor`, `utp`.
- A package README is a one-line pointer at the protocol's astral-docs spec
  dir (`api/gateway`, `api/ip`, `api/objects`); `api/apphost` extends the
  form with topic links. Only `nat`, `tor`, and `utp` lack one — their
  protocols have no astral-docs spec to point at yet.

## pub.go

- `pub.go` at the repo root is package `pub`, so a blank import of the
  module path `github.com/cryptopunkscc/astral-go` imports it.
- It blank-imports `astral`, `astral/log`, and all 20 `api` packages; their
  `init` functions run `astral.Add`, so the one import registers the full
  wire surface's blueprints.
- Extend `pub.go` when a new package registers blueprints (`.ai/rules.md`
  "Layout"); a package left out registers its types only when imported
  directly.
