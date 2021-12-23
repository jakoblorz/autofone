package endpointrequirements

import (
	"github.com/jakoblorz/metrikxd/ent"
	"github.com/jakoblorz/metrikxd/pkg"
)

type State struct {
	ent.Endpoint
	pkg.SubscriptionReceiver
}
