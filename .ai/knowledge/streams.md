# streams/

Dependency-free helpers over stdlib `io`. The package imports only the
standard library and nothing from this module, so it sits below every other
package and applies to any `io.Reader`/`io.Writer`.

## Composing Streams

* `Pipe()` returns two `PipeEnd` values forming an in-memory bidirectional
  pipe; each end is an `io.ReadWriteCloser`. A read or write error on one end
  closes the opposite direction, so the peer sees EOF. Use it as an in-process
  stand-in for a connection.
* `Join(left, right io.ReadWriteCloser)` copies between two streams in both
  directions until either reaches EOF or an error. It closes both streams and
  returns bytes written each way plus the first non-EOF error. Use it to
  splice two live connections together (relay, proxy).
* `ReadWriteCloseSplit{Reader, Writer, Closer}` composes independent
  `io.Reader`, `io.Writer`, and `io.Closer` values into one
  `io.ReadWriteCloser`.

## Reading

* `ContextReader` makes reads cancelable: `ReadContext(ctx, p)` stops waiting
  when `ctx` ends while the underlying `Read` continues in the background and
  buffers its result for the next call (chunks of `ReadBufferSize` = 16 KiB).
  `WithContext(ctx)` adapts it back to a plain `io.Reader`.
* `LimitedReader{ReadCloser, Limit}` returns `io.EOF` after exactly `Limit`
  bytes; unlike `io.LimitedReader` it propagates `Close`.
* `Skip(r, n)` reads and discards exactly `n` bytes; returns the first read
  error.
* `ReadCounter` wraps a reader and counts bytes read; `Total()` reports the
  count.

## Writing

* `AsyncWriter` decouples writers from a slow sink: writes are buffered in
  memory and flushed to the underlying `io.WriteCloser` by a background
  goroutine. `Write` never blocks; it fails with `ErrBufferOverflow` when the
  buffer is full. The caller owns the passed slice until the `AfterFlush`
  callback fires. `Sync` blocks until the buffer drains; `Close` only signals
  the flusher — wait on `Done()`, then check `Err()`.
* `Dispenser` is a credit-gated `io.WriteCloser`: `Write` sends at most the
  current limit and blocks until `Increase(i)` grants more credit or `Close`
  aborts it. `SetUnlimited(true)` bypasses the limit; `SetOutput` swaps the
  target, also mid-wait; `Flush` blocks until the credit is spent.
* `WriteCounter` wraps a writer and counts bytes written;
  `NewWriteCounter(nil)` discards output and only counts.
* `NilWriter`, `NilCloser`, `NilWriteCloser` are no-op sinks; `NilWriter`
  reports the full length as written.

## Wire Sequences

* `ReadAllFrom(r, v ...io.ReaderFrom)` and `WriteAllTo(w, v ...io.WriterTo)`
  run each element's `ReadFrom`/`WriteTo` in order, stop at the first error,
  and return the byte total. Use them to decode or encode a fixed sequence of
  primitives or wire objects.
