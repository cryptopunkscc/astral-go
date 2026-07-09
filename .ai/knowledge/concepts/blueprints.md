# Blueprints

The `Blueprints` registry (`astral/blueprints.go`) maps object type names to
prototypes so decoders can materialize objects by name. Kinds, field specs,
name rules, and limits are specified in
[blueprints](../../system/topics/blueprints.md).

## Registry

* Compile-time prototypes (registered with `astral.Add`) and runtime
  `*astral.Blueprint` values (struct kind and alias kind) share one map
  (`Blueprints.entries`); each name maps to exactly one entry.
* A registry can have a `Parent`; lookups walk the chain, and registration
  refuses names already present anywhere in the chain.
* `astral.Add` registers compile-time prototypes only; it rejects a populated
  `*Blueprint` (use `Register`/`RegisterBlueprint` for runtime registration).
* `RegisterBlueprint` validates the Blueprint, requires referenced types
  (struct kind) or the underlying primitive (alias kind) to be already
  registered, stores a defensive clone, and returns the descriptor's
  content-addressed `ObjectID`. Any name collision — against a prototype or a
  prior runtime Blueprint — returns `astral.ErrAlreadyRegistered`.
* `DefaultBlueprints()` is the process-wide registry; package-level
  `astral.Add`, `astral.New`, and `astral.Register` target it.

## Materialization

* `New(typeName)` returns a zero-value object, or nil when the name is
  unregistered.
* A compile-time prototype materializes as its typed Go value; a runtime
  Blueprint materializes as `*astral.RuntimeObject`.

## Aliases

* A compile-time prototype declares itself a primitive alias by implementing
  `astral.PrimitiveAlias` (`UnderlyingPrimitive() string`); `astral.Add` is
  then enough.
* `BlueprintOf(v)` derives a `*Blueprint` from the Go type by reflection —
  alias kind for `PrimitiveAlias` types — without storing it, so local `New`
  keeps returning the typed Go value while remote peers materialize a
  `*RuntimeObject`.

## Sync Order

* `OrderedBlueprints()` returns all registered names: compile-time prototypes
  (alpha-sorted), then aliases (alpha-sorted), then runtime Blueprints
  topo-sorted by reference.
* `AllBlueprints()` returns `*Blueprint` values for sync replay: aliases
  first, then Blueprints derived from struct prototypes, then runtime struct
  Blueprints topo-sorted by reference. Per-entry derivation failures are
  aggregated into the returned error; the slice keeps the successful entries.
* In both orders aliases precede struct Blueprints so a `RefSpec` to an alias
  resolves when a peer replays the sequence.

## Client Sync

* `apps.WithBlueprintSync()` (`lib/apps/blueprint_sync.go`) installs
  `SyncBlueprints` as a registration hook, so the app's Blueprints are pushed
  on every (re)connect.
* `SyncBlueprints` builds the local Blueprint list once per process
  (`sync.Once` over `DefaultBlueprints().AllBlueprints()`), lists the node's
  registered names, and pushes only the missing Blueprints in
  `AllBlueprints` order.
* The node answers a name collision with `ErrAlreadyRegistered` as a wire
  error; the client wrapper (`api/objects/client/register.go`) restores the
  sentinel so `errors.Is(err, astral.ErrAlreadyRegistered)` works.
* Op semantics:
  [objects.blueprints](../../system/protocols/objects/ops/objects.blueprints.md),
  [objects.register_blueprint](../../system/protocols/objects/ops/objects.register_blueprint.md).
