package runtime

import (
	"github.com/hellcats88/abstracte/env"
	"github.com/hellcats88/abstracte/logging"
	"github.com/hellcats88/abstracte/storage"
	"github.com/hellcats88/abstracte/tenant"
)

// Context defines a runtime group of common information
type Context struct {
	log    logging.Context
	tx     storage.Transaction
	tenant tenant.Context
	env    env.Context
}

// New creates an instance of runtime context with the associated logging info
func New(log logging.Context, tx storage.Transaction, tenant tenant.Context) Context {
	return Context{
		log:    log,
		tx:     tx,
		tenant: tenant,
		env:    env.New("Global"),
	}
}

// New creates an instance of runtime context with the associated logging info for a specific environment
func NewWithEnv(log logging.Context, tx storage.Transaction, tenant tenant.Context, env env.Context) Context {
	return Context{
		log:    log,
		tx:     tx,
		tenant: tenant,
		env:    env,
	}
}

// New clones a context using new transaction
func NewFromTx(from Context, tx storage.Transaction) Context {
	return Context{
		log:    from.log,
		env:    from.env,
		tx:     tx,
		tenant: from.tenant,
	}
}

// Log return the current Logging Context
func (c Context) Log() logging.Context {
	return c.log
}

// Tx return the current Transaction Context
func (c Context) Tx() storage.Transaction {
	return c.tx
}

// Tenant return the current Tenant Context
func (c Context) Tenant() tenant.Context {
	return c.tenant
}

// Env return the current environment Context
func (c Context) Env() env.Context {
	return c.env
}
