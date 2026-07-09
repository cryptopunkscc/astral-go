# Auth

`api/auth` defines the authorization wire surface: typed action objects and
the contract types that grant one identity permissions on another's behalf.
Authorization dispatch (handlers, the contract delegation-chain walk) lives in
astrald mod/auth.

## Actions

* Actions are typed objects embedding `auth.Action` (`Nonce` + `ActorID`):
  see [mod.auth.action](../../system/protocols/auth/types/mod.auth.action.md).
* The actor lives inside the action object, not as a separate argument:
  `Actor() *Identity`, `SetActor(*Identity)`.
* `NewAction(actor)` returns an `Action` with a fresh nonce.
* `ActionObject` interface: `astral.Object`, `Id() Nonce`, `Actor()`,
  `SetActor()`.

## Constraints

* `Constrainable` is an optional interface on actions:
  `ApplyConstraints(*astral.Bundle) bool`.
* `Permit.Allows` consults it; an action that does not implement
  `Constrainable` is permitted unconditionally by any permit naming its type.

## Contracts

`Contract` (`api/auth/contract.go`):

* Fields and grant semantics: see
  [mod.auth.contract](../../system/protocols/auth/types/mod.auth.contract.md).
* Implements `crypto.SignableTextObject`: `SignableHash` is the contract's
  `ObjectID.Hash`; `SignableText` is the human-readable grant string.
* `HasPermit(actionType)` returns all permits matching the action type; empty
  means the contract grants no such permission.

`Permit`:

* `Action` (`String8` action object type) and optional `Constraints`
  (`*astral.Bundle`): see
  [mod.auth.permit](../../system/protocols/auth/types/mod.auth.permit.md).
* `Delegation` (`astral.Uint8`) is the number of hops allowed below a link
  carrying this permit; 0 = non-delegable. The field is not yet in the permit
  spec page.
* `Allows(action)` matches the action type, then evaluates constraints for
  `Constrainable` actions.

`SignedContract`:

* Wraps `*Contract` with `IssuerSig` and `SubjectSig` (`*crypto.Signature`);
  either may be nil before the signing step completes. Fields: see
  [mod.auth.signed_contract](../../system/protocols/auth/types/mod.auth.signed_contract.md).
* `IsNil()` is true for a nil receiver or an embedded nil `*Contract`.

## Sentinels And Ops

* Sentinels: `ErrAlreadySigned`, `ErrInvalidContract`, `ErrContractExpired`
  (`api/auth/errors.go`).
* Op names: `MethodIndex` (`auth.index`), `MethodSignContract`
  (`auth.sign_contract`); `api/auth/client` wraps both (`IndexContract`,
  `SignContract`).
* `Contract`, `Permit`, and `SignedContract` register with `astral.Add` in
  `init`.
