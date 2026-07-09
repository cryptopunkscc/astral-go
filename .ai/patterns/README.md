# Patterns

Implementation recipes from astral-go source.

Load only files that match the task.

| Task / Keywords | Read |
|---|---|
| protocol client, `api/<p>/client/`, Client, New, Default, SetDefault, queryCh, op wrapper, one op per file, channel.Expect, PassErrors | `operations.md` |
| RouteQuery, Router, RouteNotFound, Reject, RejectWithCode, Accept, PriorityRouter, custom router | `routing.md` |
| wire object, typed payload, Objectify, ObjectType, WriteTo, ReadFrom, astral.Add, MarshalJSON, pub.go | `objects.md` |
| mutex, context-aware channel, sig.RecvErr, sig.Recv, sig.Send, sig.Map, sig.Set, sig.Queue, sig.Subscribe, sig.Ring, sig.Pool | `concurrency.md` |
| add an example, runnable program | `../../examples/README.md` |

## Boundary

- File references in pattern docs are authoritative.
- Verify against source before copying a pattern into new code.
