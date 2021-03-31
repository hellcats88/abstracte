package api

type C struct {
	Handler  interface{}
	Priority int
}

type CArr []C

func (c CArr) Len() int {
	return len(c)
}

func (c CArr) Less(i, j int) bool {
	return c[i].Priority < c[j].Priority
}

func (c CArr) Swap(i, j int) {
	tmp := c[i]
	c[i] = c[j]
	c[j] = tmp
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

type Config struct {
	Log         ConfigLog
	Tenant      ConfigTenant
	Tx          ConfigTx
	InputModel  interface{}
	InputParams []string
	QueryParams interface{}
	Headers     interface{}
}

type ConfigGroup struct {
	Config
	Handlers CArr
}

type ConfigRoute struct {
	Config
	CustomHandlers CArr
}
