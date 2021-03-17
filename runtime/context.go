package runtime

import (
	"github.com/hellcats88/abstracte/env"
	"github.com/hellcats88/abstracte/logging"
	"github.com/hellcats88/abstracte/storage"
	"github.com/hellcats88/abstracte/tenant"
)

// Context defines a runtime group of common information
type Context interface {
	Log() logging.Context
	Tx() storage.Transaction
	Tenant() tenant.Context
	Env() env.Context
}
