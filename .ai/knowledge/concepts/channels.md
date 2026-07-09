# Channels

A `Channel` (`astral/channel`) carries typed `Object` values over a raw stream.

* Choose the encoding when constructing the channel: `bin` (the default),
  `json`, `text`, or `canonical`; senders also accept `base64` and `render`.
  Encoding semantics are specified in
  [channel](../../system/core-definitions/channel.md).
* `EOS` marks the end of the stream.
* Receive loops stop on transport `io.EOF`, an error, or a helper condition.

## Receive Styles

| API       | Use when                                                                    |
|-----------|-----------------------------------------------------------------------------|
| `Switch`  | Client calls. Type-dispatch loop; stops on EOF, error, or helper condition. |
| `Collect` | Op handlers. Runs a callback per object; caller type-switches.              |
| `Handle`  | Subscriptions. Runs a context-cancellable loop.                             |

## Switch Helpers

| Helper                       | Effect                                       |
|------------------------------|----------------------------------------------|
| `channel.Expect(&ptr)`       | receive one T, stop                          |
| `channel.ExpectString(&s)`   | receive one T, store `T.String()`, stop      |
| `channel.PassErrors`         | `*astral.ErrorMessage` → Go error            |
| `channel.BreakOnEOS`         | stop on EOS, return nil                      |
| `channel.ExpectAck`          | receive Ack, stop                            |
| `channel.Collect[T](&slice)` | append each T; pair with `BreakOnEOS`        |
| `channel.Chan[T](ch)`        | forward T into a Go channel; pair with `BreakOnEOS` |
| `channel.WithContext(ctx)`   | cancel Switch on ctx                         |
| `channel.WithTimeout(d)`     | cancel Switch after duration                 |

## Sending

* End every stream with `EOS`.
* Send mid-stream errors as `astral.Err(err)`.
* Do not signal mid-stream errors by closing the channel.

## Locked Writes

* `channel.WithLockedWrites()` wraps `Send` in a mutex so concurrent senders
  do not interleave object bytes on the underlying writer.
* Use it when multiple goroutines share one `Channel` over a single transport.
* Individual `Sender` constructors do not honour this option; pass it to
  `channel.New`.
