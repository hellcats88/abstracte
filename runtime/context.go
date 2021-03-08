package runtime

import (
	"github.com/hellcats88/abstracte/env"
	"github.com/hellcats88/abstracte/logging"
	"github.com/hellcats88/abstracte/storage"
	"github.com/hellcats88/abstracte/tenant"
)

// Context defines a runtime group of common information
type Context struct {
	Log    logging.Context
	Tx     storage.Transaction
	Tenant tenant.Context
	Env    env.Context
}

// New creates an instance of runtime context with the associated logging info
func New(log logging.Context, tx storage.Transaction, tenant tenant.Context) Context {
	return Context{
		Log:    log,
		Tx:     tx,
		Tenant: tenant,
		Env:    env.New("Global"),
	}
}

// New creates an instance of runtime context with the associated logging info for a specific environment
func NewWithEnv(log logging.Context, tx storage.Transaction, tenant tenant.Context, env env.Context) Context {
	return Context{
		Log:    log,
		Tx:     tx,
		Tenant: tenant,
		Env:    env,
	}
}

// New clones a context using new transaction
func NewFromTx(from Context, tx storage.Transaction) Context {
	return Context{
		Log:    from.Log,
		Env:    from.Env,
		Tx:     tx,
		Tenant: from.Tenant,
	}
}
