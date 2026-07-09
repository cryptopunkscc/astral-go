# AI Workspace

Vendor-neutral AI context for astral-go, the Go SDK of the Astral Network.

astral-go holds the network's app-facing surface: the astral primitives, the
protocol wire types and clients under `api/`, and the client libraries under
`lib/`. Apps import this module. astrald (the node daemon) imports this module.
This module imports neither.

## Load Order

1. `.ai/README.md` — this file: orientation and the index below.
2. `.ai/rules.md` — always-on standards.

Then load a scoped note from the index only when it matches the task. Protocol
meaning — encodings, op semantics, wire types — lives in the spec (`.ai/system/`,
the astral-docs submodule); notes here cover the Go SDK and link the spec,
never restate it.

## Authority

1. User instruction
2. Code/tests
3. `.ai/system/` (the astral-docs spec, pinned as a submodule)
4. `.ai/rules.md`
5. `.ai/knowledge/`
6. `.ai/patterns/`

Call out conflicts.

## Knowledge — how the SDK works

| Keywords | Read |
|---|---|
| Query, InFlightQuery, Launch, Extra, OriginLocal, OriginNetwork, Router, RouteQuery, ErrRouteNotFound, HasRoutingPriority, PriorityRouter | `knowledge/query.md` |
| Auth, Action, ActionObject, NewAction, Constrainable, Contract, Permit, SignedContract, HasPermit, Allows, Delegation, MethodSignContract, ErrContractExpired | `knowledge/auth.md` |
| lib/astrald, lib/apphost, lib/ipc, lib/routing, lib/apps, lib/query, QueryChannel, WithTarget, OpRouter, ScopeRouter, PriorityRouter, IncomingQuery, Serve, AppRegistrar, client library | `knowledge/lib.md` |
| Op, op handler, routing.NewOp, handler signature, args struct, query.Parse, query tag, required, key, skip, AddOp, AddStructPrefix, Spec, snake_case, ErrInvalidSignature | `knowledge/operations.md` |
| Channel, Switch, Collect, Handle, EOS, channel.Expect, ExpectAck, PassErrors, BreakOnEOS, WithContext, WithTimeout, WithLockedWrites, astral.Err, encoding bin json text canonical | `knowledge/channels.md` |
| Serialization, wire format, Objectify, ObjectType, canonical encoding, binary framing, Stamp, ObjectID hashing, supported field kinds, Stringify, astral primitives | `knowledge/wire.md` |
| tree.Value, typed cell, Bind, BindPath, Get, Follow, Set, Clear, NoLocal, astral.Nil, live binding | `knowledge/tree.md` |
| Blueprints, Blueprint, astral.Add, astral.Register, RegisterBlueprint, astral.New, PrimitiveAlias, BlueprintOf, RuntimeObject, OrderedBlueprints, AllBlueprints, SyncBlueprints, WithBlueprintSync, ErrAlreadyRegistered | `knowledge/blueprints.md` |
| streams/, Pipe, Join, ReadWriteCloseSplit, ContextReader, LimitedReader, Skip, ReadCounter, AsyncWriter, Dispenser, WriteCounter, ReadAllFrom, WriteAllTo | `knowledge/streams.md` |
| astral/log, Logger, Log, Logv, Entry, Tag, SetPrefix, SetFilter, EntryLogger, ToSnakeCase, styles, theme, views | `knowledge/log.md` |
| api/ conventions, package shape, module.go, Method* constants, client/ packages, package README, pub.go, blank import, registration aggregator | `knowledge/api.md` |

Per-protocol wire types and op semantics are in the spec
(`.ai/system/protocols/<p>/`, pointed at by each `api/<p>/README.md`); the
client wrappers are one method per file under `api/<p>/client/`.

## Patterns — how to write it

| Task / Keywords | Read |
|---|---|
| protocol client, `api/<p>/client/`, Client, New, Default, SetDefault, queryCh, op wrapper, one op per file, channel.Expect, PassErrors | `patterns/client.md` |
| RouteQuery, Router, RouteNotFound, Reject, RejectWithCode, Accept, PriorityRouter, custom router | `patterns/routing.md` |
| wire object, typed payload, Objectify, ObjectType, WriteTo, ReadFrom, astral.Add, MarshalJSON, pub.go | `patterns/wire-type.md` |
| mutex, context-aware channel, sig.RecvErr, sig.Recv, sig.Send, sig.Map, sig.Set, sig.Queue, sig.Subscribe, sig.Ring, sig.Pool | `patterns/concurrency.md` |
| add an example, runnable program | `../examples/README.md` |

File references in the notes are authoritative; verify against source before
copying a pattern into new code.
