# astrald

A client library for `astrald`. The default client routes queries through the
local node's apphost service (`lib/apphost`), using the access token from
`ASTRALD_APPHOST_TOKEN`.

## Quick start

### Sending a query

```go
package main

import (
	"fmt"

	dirclient "github.com/cryptopunkscc/astral-go/api/dir/client"
	"github.com/cryptopunkscc/astral-go/lib/astrald"
	"github.com/cryptopunkscc/astral-go/lib/query"
)

func main() {
	// create a new context
	ctx := astrald.NewContext()

	fmt.Printf("authenticated as %s on %s\n",
		astrald.GuestID(),
		astrald.HostID(),
	)

	// query the local node
	conn, err := astrald.Query(ctx, "method", query.Args{"arg": "val"})

	// ... or use a helper method to get a typed object channel
	ch, err := astrald.QueryChannel(ctx, "method", query.Args{"arg": "val"})

	// ... or query some other node
	nodeID, err := dirclient.ResolveIdentity(ctx, "nodealias")

	conn, err = astrald.WithTarget(nodeID).Query(ctx, "method", nil)

	// a Conn is a raw byte stream; close it when done
	defer conn.Close()

	// read the reply from the channel
	reply, err := ch.Receive()
	if err != nil {
		return
	}
	fmt.Println(reply)
}
```

### Listening for queries

Serving queries lives in `lib/apps`: pass a router to `apps.Serve` and it
keeps the app registered with the node until the context ends.

```go
package main

import (
	"github.com/cryptopunkscc/astral-go/astral"
	"github.com/cryptopunkscc/astral-go/lib/apps"
	"github.com/cryptopunkscc/astral-go/lib/astrald"
	"github.com/cryptopunkscc/astral-go/lib/routing"
)

type API struct{}

type helloArgs struct {
	Name string `query:"required"`
}

// Hello answers queries of the form "hello?name=alice".
func (api *API) Hello(ctx *astral.Context, q *routing.IncomingQuery, args helloArgs) error {
	ch := q.Accept()
	defer ch.Close()

	return ch.Send(astral.NewString16("hello, " + args.Name))
}

func main() {
	if err := apps.Serve(astrald.NewContext(), routing.NewApp(&API{})); err != nil {
		panic(err)
	}
}
```

See `lib/apps` for registration options and the generated `objects.*` routes.

### Resolving aliases

```go
package main

import (
	"fmt"
	"os"

	dirclient "github.com/cryptopunkscc/astral-go/api/dir/client"
	"github.com/cryptopunkscc/astral-go/lib/astrald"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: astral-resolve <alias>")
		return
	}

	ctx := astrald.NewContext()

	identity, err := dirclient.ResolveIdentity(ctx, os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}

	fmt.Println(identity)
}
```
