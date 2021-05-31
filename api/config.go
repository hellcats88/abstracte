package api

type C struct {
	Handler interface{}
}

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
const InputParamsKey = "_abstracte_api_inputparams_key"
const QueryParamsKey = "_abstracte_api_queryparams_key"
const HeadersModelKey = "_abstracte_api_headersmodel_key"

type Config interface {
	Valid() bool
}

type ConfigBuilder interface {
	Log(ConfigLog) ConfigBuilder
	CustomLog(C) ConfigBuilder
	Tenant(ConfigTenant) ConfigBuilder
	CustomTenant(C) ConfigBuilder
	Tx(ConfigTx) ConfigBuilder
	CustomTx(C) ConfigBuilder
	Headers(interface{}) ConfigBuilder
	CustomHeaders(C) ConfigBuilder
	InputModel(interface{}) ConfigBuilder
	CustomInputModel(C) ConfigBuilder
	InputParams(name []string) ConfigBuilder
	CustomInputParam(C) ConfigBuilder
	QueryParams(interface{}) ConfigBuilder
	CustomQueryParams(C) ConfigBuilder
	Custom(C) ConfigBuilder
	Build() Config
}
