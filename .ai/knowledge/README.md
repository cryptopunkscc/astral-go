# Knowledge Index

Repo implementation knowledge. If this conflicts with code, trust code and update this.

## Concepts

Concept pages explain cross-package SDK ideas.

| Keywords | Read |
|---|---|
| Query, InFlightQuery, Launch, Extra, OriginLocal, OriginNetwork, Router, RouteQuery, ErrRouteNotFound, HasRoutingPriority, PriorityRouter | `concepts/query.md` |
| Auth, Action, ActionObject, NewAction, Constrainable, Contract, Permit, SignedContract, HasPermit, Allows, Delegation, MethodSignContract, ErrContractExpired | `concepts/auth.md` |
| lib/astrald, lib/apphost, lib/ipc, lib/routing, lib/apps, lib/query, QueryChannel, WithTarget, OpRouter, ScopeRouter, PriorityRouter, IncomingQuery, Serve, AppRegistrar, client library | `concepts/lib.md` |
| Op, op handler, routing.NewOp, handler signature, args struct, query.Parse, query tag, required, key, skip, AddOp, AddStructPrefix, Spec, snake_case, ErrInvalidSignature | `concepts/operations.md` |
| Channel, Switch, Collect, Handle, EOS, channel.Expect, ExpectAck, PassErrors, BreakOnEOS, WithContext, WithTimeout, WithLockedWrites, astral.Err, encoding bin json text canonical | `concepts/channels.md` |
| Serialization, wire format, Objectify, ObjectType, canonical encoding, binary framing, Stamp, ObjectID hashing, supported field kinds, Stringify, astral primitives | `concepts/wire.md` |
| tree.Value, typed cell, Bind, BindPath, Get, Follow, Set, Clear, NoLocal, astral.Nil, live binding | `concepts/tree.md` |
| Blueprints, Blueprint, astral.Add, astral.Register, RegisterBlueprint, astral.New, PrimitiveAlias, BlueprintOf, RuntimeObject, OrderedBlueprints, AllBlueprints, SyncBlueprints, WithBlueprintSync, ErrAlreadyRegistered | `concepts/blueprints.md` |

## API

`api.md` covers the conventions of the `api/` tree. Per-protocol wire types
and op semantics live in the astral-docs spec (`.ai/system/protocols/<p>/`,
pointed at by each `api/<p>/README.md`); the client wrappers are one method
per file under `api/<p>/client/`.

| Keywords | Read |
|---|---|
| `api/` conventions, package shape, module.go, Method* constants, client/ packages, package README, pub.go, blank import, registration aggregator | `api.md` |

## SDK Notes

| Keywords | Read |
|---|---|
| `streams/`, Pipe, Join, ReadWriteCloseSplit, ContextReader, LimitedReader, Skip, ReadCounter, AsyncWriter, Dispenser, WriteCounter, ReadAllFrom, WriteAllTo | `streams.md` |
| `astral/log/`, Logger, Log, Logv, Entry, Tag, SetPrefix, SetFilter, EntryLogger, ToSnakeCase, styles, theme, views | `astral/log.md` |

## Rules and Patterns

| Keywords | Read |
|---|---|
| Coding rule, constraint, invariant, boundary, layout, wire-type rules, astral.Adapt, naming, style | `../rules.md` |
| Pattern, recipe, skeleton, boilerplate, how to write, client wrapper, op handler, example | `../patterns/README.md` |
