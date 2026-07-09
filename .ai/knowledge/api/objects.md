# api/objects

Wire surface of the object layer: wire types, op-name constants, repository
names, the app-facing extension interfaces, the `Writer` contract, and the
remote client. The store itself — repositories, tracking, purge, extension
dispatch — is astrald mod/objects; op semantics are specified in
[protocols/objects](../../system/protocols/objects/README.md).

## Wire surface

- `module.go` declares `ModuleName`, the `Method*` op-name constants
  (`objects.create`, `objects.read`, `objects.search`, …), and the `Repo*`
  names of the default repositories and groups (`main`, `device`, `memory`,
  `local`, `removable`, `virtual`, `network`, `system`).
- Wire types register with `astral.Add` in their defining files:
  `Descriptor`, `SearchQuery`, `QueryTag`, `SearchResult`, `Probe`,
  `RepositoryInfo`, `CommitMsg`, `ReadObjectAction`, `CreateObjectAction`.
  Encodings are specified under
  [protocols/objects/types](../../system/protocols/objects/types/).
- `ErrTagUnsupported` (`err_tag_unsupported.go`) signals a search tag the
  describer does not support; matchable with `errors.Is`.

## Extension interfaces

Provider contracts apps implement; registration, fan-out, and dedup dispatch
are node-side in astrald mod/objects.

- `Describer` (`module.go`): `DescribeObject(ctx, id) (<-chan *Descriptor, error)`.
- `Searcher` (`search.go`): `SearchObject(ctx, SearchQuery) (<-chan *SearchResult, error)`.
- `SearchPreprocessor` (`search.go`): `PreprocessSearch(*Search)` mutates a
  `Search` in place before it runs.
- `Finder` (`find.go`): `FindObject(ctx, id) (<-chan *astral.Identity, error)`.

## Writer

- `objects.Writer` (`writer.go`): `Write`, `Commit`, `Discard`. Every writer
  obtained from `Create` must end in exactly one `Commit` or `Discard`.
- The client's `Create(ctx, repo, alloc)` opens `objects.create` (optional
  `repo` and `alloc` args), waits for `Ack`, and returns a `Writer` over the
  open channel.
- `Write` sends chunks as `*astral.Blob`; `Commit` sends `CommitMsg` and
  receives the `*astral.ObjectID`; `Discard` closes the channel without
  committing.

## Client

`Client` (`api/objects/client/client.go`): `New(targetID, astralClient)`
targets any node, a nil client falls back to `astrald.Default()`; most
wrappers also exist as package-level functions on `Default()`.

- Storage: `Create`, `Store`, `Read`, `Delete`, `Purge`, `NewMem`,
  `Repositories`.
- Discovery: `Scan`, `Search`, `Describe`, `Find`, `Probe`, `GetType`.
- Delivery: `Push` sends one object and expects a boolean answer (false
  fails with "rejected"); `Echo` opens an `objects.echo` channel the caller
  drives and must close.
- Streaming wrappers (`Scan`, `Search`, `Describe`, `Find`, `Purge`) return
  `(<-chan T, *error)`; the error pointer is valid once the channel closes.
  `Scan(follow=true)` emits one nil ID between snapshot and live updates.
- `RegisterDescriber`/`RegisterFinder`/`RegisterSearcher` register the caller
  as a remote provider (ops `objects.register_*`) and block until `Ack`.

## Blueprint registration

- `Blueprints(ctx)` lists the type names registered on the target in
  dependency order (`objects.blueprints`, `String8` stream until `EOS`).
- `Register(ctx, o)` (`register.go`) pushes a runtime `*astral.Blueprint`
  (struct or alias kind) over `objects.register_blueprint` and returns its
  content-addressed `ObjectID`. A name collision arrives as a wire error
  string; `Register` matches the `astral.ErrAlreadyRegistered` prefix and
  returns the in-process sentinel, so `errors.Is(err,
  astral.ErrAlreadyRegistered)` works across the wire.
- Registry semantics and sync live in
  [concepts/blueprints](../concepts/blueprints.md).
