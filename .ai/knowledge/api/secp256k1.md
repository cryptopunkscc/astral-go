# api/secp256k1

Key helpers and constants of the `secp256k1` protocol: private-key
generation, public-key derivation, and the bridge between `astral.Identity`
and `crypto.PublicKey`. `client/` is the protocol's RPC client. Protocol
meaning lives in the spec:
[protocols/secp256k1](../../system/protocols/secp256k1/).

## Constants

`module.go`: `KeyType` = `secp256k1` — the canonical key-type string
consumed across the crypto stack (astrald mod/crypto, mod/bip137sig,
mod/coldcard) — and `MethodNew` = `secp256k1.new`.

## Helpers

All in `module.go`, operating on `api/crypto` key types:

- `New()` — generates a fresh `*crypto.PrivateKey` with `Type = KeyType`
  and the serialized private key.
- `PublicKey(key)` — derives the public key via
  `secp256k1.PrivKeyFromBytes(...).PubKey().SerializeCompressed()`; nil
  when `key.Type != KeyType`.
- `FromIdentity(identity)` — `*astral.Identity` to `*crypto.PublicKey`:
  wraps `identity.PublicKey().SerializeCompressed()`.
- `Identity(key)` — `*crypto.PublicKey` back to `*astral.Identity` via
  `ParsePubKey` and `astral.IdentityFromPubKey`; nil on parse failure.

## Invariants

- `FromIdentity` and `PublicKey` emit 33-byte compressed public keys; the
  compressed form is the standard serialization.
- The helpers return nil on failure instead of an error.

## Client

`client/` wraps a `lib/astrald` client with an optional target identity
(`New(targetID, client)`, `Default()`). One wrapper:
`Client.NewKey(ctx)` queries `secp256k1.new` and returns a fresh
`*crypto.PrivateKey`; package-level `NewKey` calls it on the default
client.

The crypto engine (public-key derivation, ASN.1 hash signing and
verification) and the `secp256k1.new` handler are node-side: astrald
mod/secp256k1.
