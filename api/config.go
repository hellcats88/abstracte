package api

type ConfigLog uint

const LogKey = "_abstracte_api_config_log_key"
const LogIdx = 0x0

const (
	ConfigLogRandom      ConfigLog = 0x0
	ConfigLogFromHeaders ConfigLog = 0x1
)

type ConfigTenant uint

const TenantKey = "_abstracte_api_config_tenant_key"
const TenantIdx = 0x1

const (
	ConfigTenantNo          ConfigTenant = 0x0
	ConfigTenantFromHeaders ConfigTenant = 0x1
)

type ConfigTx uint

const TxKey = "_abstracte_api_config_tx_key"
const TxIdx = 0x1

const (
	ConfigTxManaged   ConfigTx = 0x1
	ConfigTxUnmanaged ConfigTx = 0x2
)

const RuntimeKey = "_abstracte_api_runtime_key"
const InputModelKey = "_abstracte_api_inputmodel_key"

type Config struct {
	Log    ConfigLog
	Tenant ConfigTenant
	Tx     ConfigTx
}

type ConfigGroup struct {
	Config
	Handlers []interface{}
}

type ConfigRoute struct {
	Config
	CustomHandlers []interface{}
	ServiceHandler []interface{}
}
