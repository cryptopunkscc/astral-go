package dir

import (
	"github.com/cryptopunkscc/astral-go/astral"
)

const ModuleName = "dir"

const (
	MethodAliasMap     = "dir.alias_map"
	MethodApplyFilters = "dir.apply_filters"
	MethodFilters      = "dir.filters"
	MethodGetAlias     = "dir.get_alias"
	MethodResolve      = "dir.resolve"
	MethodSetAlias     = "dir.set_alias"
)

// Resolver is implemented by any source that can map names to identities or supply display names.
type Resolver interface {
	ResolveIdentity(string) (*astral.Identity, error)
	DisplayName(*astral.Identity) string
}
