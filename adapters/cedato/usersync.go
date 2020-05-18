package cedato

import (
	"text/template"

	"github.com/prebid/prebid-server/adapters"
	"github.com/prebid/prebid-server/usersync"
)

// TBD
func NewCedatoSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("cedato", 672, temp, adapters.SyncTypeRedirect)
}
