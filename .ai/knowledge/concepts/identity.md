# Identity

`astral.Identity` is a compressed secp256k1 public key (`astral/identity.go`);
its object type is `identity`. Encodings are specified in
[identity](../../system/primitive-types/identity.md); the roles an identity
can play (Node, App, User) in
[Identity](../../system/core-definitions/identity.md).

## Representation

* Binary form: 33 bytes; the zero identity encodes as 33 zero bytes.
* String form (`String`, `MarshalText`): 66 hex characters; the zero identity
  emits the all-zero string.
* `ParseIdentity` accepts a 66-char hex key, the all-zero hex string, or the
  string `"anyone"`; any other length fails with `ErrInvalidKeyLength`.
* JSON form: the hex string; the zero identity marshals as `"anyone"`, and
  `"anyone"` unmarshals back to it.

## Anyone

* `Anyone` is the zero identity: anonymous or wildcard.
* `IsZero()` is true for nil receivers and for empty public keys.
* `IsEqual` treats two zero identities as equal; a zero identity never equals
  a non-zero one.

Identity as network address — directory resolution and user swarms — is
node-side: astrald mod/dir and mod/user.
