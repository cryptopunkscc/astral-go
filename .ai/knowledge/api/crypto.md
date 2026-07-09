# api/crypto

Crypto object types, scheme constants, and op-name constants of the `crypto`
protocol. `client/` is a minimal RPC client. Encodings live in the spec:
[protocols/crypto/types](../../system/protocols/crypto/types/); op semantics
in [protocols/crypto/ops](../../system/protocols/crypto/ops/).

## Object types

| Type | Fields | Object type |
|---|---|---|
| `PrivateKey` | `Type String8`, `Key Bytes16` | `mod.crypto.private_key` |
| `PublicKey` | `Type String8`, `Key Bytes16` | `mod.crypto.public_key` |
| `Signature` | `Scheme String8`, `Data Bytes16` | `mod.crypto.signature` |
| `Hash` | `[]byte` | `mod.crypto.hash` |

All four register with `astral.Add` and carry JSON and text codecs.

## Signable contracts

- `SignableObject` (`signable_object.go`) — `astral.Object` +
  `SignableHash() []byte`; base interface for all contracts.
- `SignableTextObject` — adds `SignableText() string`, a human-readable
  contract text no longer than 200 characters.
- `SignableHash` must yield at least 15 bytes: the astrald mod/crypto text
  signer embeds `base64(hash[0:15])` as the commitment in the signed text.

## Constants

- Schemes (`engine.go`): `SchemeASN1` = `asn1` (default for hash
  signatures), `SchemeBIP137` = `bip137` (default for text signatures).
- Op names (`module.go`): `MethodPublicKey`, `MethodSignHash`,
  `MethodSignText`, `MethodVerifyHashSignature`, `MethodVerifyTextSignature`
  (`crypto.<op>`).

## Client

`client/` wraps a `lib/astrald` client with an optional target identity
(`New(targetID, astral)`, `Default()`). It wraps one op:
`Client.PublicKey(ctx, privateKey)` sends a `PrivateKey` on the
`crypto.public_key` channel and expects a `PublicKey` back. The signing and
verification ops have no wrappers.

Engines, key indexing, and capability dispatch are node-side: astrald
mod/crypto.
