# AI Workspace

Vendor-neutral AI context for astral-go, the Go SDK of the Astral Network.

astral-go holds the network's app-facing surface: the astral primitives, the
protocol wire types and clients under `api/`, and the client libraries under
`lib/`. Apps import this module. astrald (the node daemon) imports this module.
This module imports neither.

## Load Order

1. `.ai/README.md`
2. `.ai/rules.md`

Then use indexes. Load scoped files only when relevant:

- `.ai/knowledge/README.md` - repo implementation
- `.ai/patterns/README.md` - code recipes
- `.ai/system/` - domain/protocol truth

## Authority

1. User instruction
2. Code/tests
3. `.ai/system/` (the astral-docs spec, pinned as a submodule)
4. `.ai/rules.md`
5. `.ai/knowledge/`
6. `.ai/patterns/`

Call out conflicts.

## Roles

- `README.md` - this file: orientation and load order
- `rules.md` - compact always-on rules
- `knowledge/` - repo implementation notes
- `patterns/` - source-grounded recipes
- `system/` - domain/protocol truth (astral-docs at a pinned commit)
