# api/bip137sig

Entropy and seed wire objects, BIP-39/BIP-32 helpers, and op-name constants
of the `bip137sig` protocol. `client/` is the protocol's RPC client.
Encodings and op semantics live in the spec:
[protocols/bip137sig](../../system/protocols/bip137sig/).

## Wire objects

- `Entropy` (`entropy.go`) — raw entropy bytes, object type
  `bip137sig.entropy`.
- `Seed` (`seed.go`) — BIP-39 seed bytes, object type `bip137sig.seed`.

Both register with `astral.Add` and carry hex text codecs. `WriteTo` and
`ReadFrom` enforce lengths on both ends: `Entropy` must be 16, 20, 24, 28,
or 32 bytes (a multiple of `EntropyStepBytes` within `MinEntropyBytes`..
`MaxEntropyBytes`), else `ErrInvalidEntropyLength`; `Seed` must be exactly
`SeedLengthBytes` = 64 bytes, else `ErrInvalidSeedLength`.

## BIP-39 helpers

`bip39.go`; the English wordlist is `bip39_wordlist.go`.

- `NewEntropy(bits)` — cryptographically random `Entropy`; `bits` must be
  128-256 in 32-bit increments.
- `EntropyToMnemonic(entropy)` — entropy to mnemonic words (11 bits per
  word, SHA-256 checksum of `ENT/32` bits).
- `MnemonicToEntropy(words)` — words back to entropy; validates word count
  (12, 15, 18, 21, or 24) and checksum.
- `MnemonicToSeed(words, passphrase)` — validates the mnemonic, then
  derives a 64-byte `Seed` with PBKDF2-HMAC-SHA512, 2048 rounds, salt
  `"mnemonic" + passphrase`. Empty passphrase is valid.

## BIP-32 path parsing

`ParseDerivationPath` (`bip32.go`) parses a BIP-32 path into child indices:
`m` or `""` yields a nil path; a leading `m/` is stripped; a trailing `'`
or `h` marks a hardened index by setting bit 31 (`0x80000000`).

## Constants

- Sizes (`bip39.go`): `MinEntropyBits` = 128, `MaxEntropyBits` = 256,
  `EntropyStepBits` = 32 (byte forms `MinEntropyBytes`/`MaxEntropyBytes`/
  `EntropyStepBytes`), `SeedLengthBytes` = 64, `DefaultEntropyBits` = 128.
- Op names (`module.go`): `MethodNewEntropy`, `MethodMnemonic`,
  `MethodSeed`, `MethodDeriveKey` (`bip137sig.<op>`).
- Sentinels (`errors.go`): `ErrInvalidMnemonic`,
  `ErrInvalidMnemonicWordCount`, `ErrInvalidEntropySize`,
  `ErrInvalidEntropyLength`, `ErrInvalidSeedLength`.

## Client

`client/` wraps a `lib/astrald` client with an optional target identity
(`New(targetID, astral)`, `Default()`). One wrapper per op:

- `NewEntropy(ctx, bits)` — `Entropy` of the requested bit length
- `EntropyToMnemonic(ctx, entropy)` — sends the `Entropy`, returns the
  mnemonic as `[]string`
- `MnemonicToSeed(ctx, mnemonic, passphrase)` — sends the joined mnemonic
  as a `String16`, returns a `Seed`
- `DeriveKey(ctx, path, seed)` — sends the `Seed` with a `path` arg,
  returns a `crypto.PrivateKey`

The crypto engine, message signer, and op handlers are node-side: astrald
mod/bip137sig.
