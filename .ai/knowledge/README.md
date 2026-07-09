# Knowledge Index

Repo implementation knowledge. If this conflicts with code, trust code and update this.

## Concepts

Concept pages explain cross-package SDK ideas.

| Keywords | Read |
|---|---|
| Identity, Anyone, ParseIdentity, IsZero, IsEqual, "anyone", 33-byte key, 66 hex, ErrInvalidKeyLength | `concepts/identity.md` |
| Zone, ZoneDevice, ZoneVirtual, ZoneNetwork, ZoneAll, Zones, zone letters, Context, WithZone, IncludeZone, ExcludeZone, LimitZone | `concepts/zone.md` |
| Query, Nonce, QueryString, NewQuery, InFlightQuery, Launch, Extra, OriginLocal, OriginNetwork, Router, RouteQuery, ErrRouteNotFound, HasRoutingPriority | `concepts/query.md` |
| Auth, Action, ActionObject, NewAction, Constrainable, Contract, Permit, SignedContract, HasPermit, Allows, Delegation, MethodSignContract, ErrContractExpired | `concepts/auth.md` |
| lib/astrald, lib/apphost, lib/ipc, lib/routing, lib/apps, lib/query, QueryChannel, WithTarget, OpRouter, ScopeRouter, PriorityRouter, IncomingQuery, Serve, AppRegistrar, client library | `concepts/lib.md` |
| Op, op handler, routing.NewOp, handler signature, args struct, query.Parse, query tag, required, key, skip, AddOp, AddStructPrefix, Spec, snake_case, ErrInvalidSignature | `concepts/operations.md` |
| Channel, Switch, Collect, Handle, EOS, channel.Expect, ExpectAck, PassErrors, BreakOnEOS, WithContext, WithTimeout, WithLockedWrites, astral.Err, encoding bin json text canonical | `concepts/channels.md` |
| Serialization, wire format, Objectify, ObjectType, canonical encoding, binary framing, Stamp, ObjectID hashing, supported field kinds, Stringify, astral primitives | `concepts/wire.md` |
| tree.Value, typed cell, Bind, BindPath, Get, Follow, Set, Clear, NoLocal, astral.Nil, live binding | `concepts/tree.md` |
| Blueprints, Blueprint, astral.Add, astral.Register, RegisterBlueprint, astral.New, PrimitiveAlias, BlueprintOf, RuntimeObject, OrderedBlueprints, AllBlueprints, SyncBlueprints, WithBlueprintSync, ErrAlreadyRegistered | `concepts/blueprints.md` |

## API

Per-protocol wire-surface notes. Read the protocol note when working under
`api/<p>/`; `api.md` covers the tree's conventions.

| Keywords | Read |
|---|---|
| `api/` conventions, package shape, module.go, Method* constants, client/ packages, package README, pub.go, blank import, registration aggregator | `api.md` |
| `api/apphost/`, guest protocol, AccessToken, HostInfoMsg, AuthTokenMsg, RouteQueryMsg, QueryAcceptedMsg, RegisterHandlerMsg, BindMsg, ErrCode, MethodCreateToken, HoldObject, apphost client | `api/apphost.md` |
| `api/bip137sig/`, Entropy, Seed, mnemonic, BIP-39, BIP-32, EntropyToMnemonic, MnemonicToSeed, ParseDerivationPath, MethodDeriveKey | `api/bip137sig.md` |
| `api/crypto/`, PrivateKey, PublicKey, Signature, Hash, SignableObject, SignableTextObject, SignableHash, SchemeASN1, SchemeBIP137, MethodSignHash, MethodSignText | `api/crypto.md` |
| `api/dir/`, identity directory, Alias, AliasMap, ResolveIdentity, GetAlias, SetAlias, Filters, ApplyFilters, MethodResolve, EnableCache | `api/dir.md` |
| `api/exonet/`, Endpoint interface, Network, Address, Pack, dialable address, EndpointWithTTL | `api/exonet.md` |
| `api/indexing/`, IndexMsg, UnindexMsg, ChangeAckMsg, RegisterIndexer, Subscribe, Subscription, Next, Ack, Fail, ErrIndexingTemporarilyFailed | `api/indexing.md` |
| `api/kcp/`, KCP, kcp.Endpoint, EndpointLocalMapping, MethodNewEphemeralListener, SetEndpointLocalPort, ListEndpointLocalMappings | `api/kcp.md` |
| `api/nat/`, hole punch, Hole, PunchSignal, ConsumeHoleSignal, Puncher, PunchProtocol, PunchResult, NodePunch, NodeConsumeHole, ListHoles | `api/nat.md` |
| `api/objects/`, Descriptor, SearchQuery, SearchResult, Probe, Finder, Searcher, Describer, Writer, Create, Commit, Discard, Push, Echo, Repo names, RegisterDescriber, register_blueprint, objects client | `api/objects.md` |
| `api/secp256k1/`, KeyType, secp256k1.new, MethodNew, NewKey, FromIdentity, compressed public key | `api/secp256k1.md` |
| `api/services/`, service discovery, Update, services.update, Discover, MethodDiscover, MethodSync, follow, snapshot | `api/services.md` |
| `api/tcp/`, TCP, tcp.Endpoint, mod.tcp.endpoint, ParseEndpoint, CreateEphemeralListener, CloseEphemeralListener | `api/tcp.md` |
| `api/tor/`, Tor, Digest, onion, tor.Endpoint, DigestFromString, base32 | `api/tor.md` |
| `api/tree/`, tree.Node, tree.Query, tree.Get, tree.Follow, NodeOps, MethodGet, MethodSet, MethodDelete, MethodList, MethodMountRemote, ErrNoValue, ErrTypeMismatch, remote tree client | `api/tree.md` |
| `api/utp/`, uTP, utp.Endpoint, mod.utp.endpoint, UDPAddr, ParseEndpoint | `api/utp.md` |

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
