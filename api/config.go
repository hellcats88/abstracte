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

type Config struct {
	Log    ConfigLog
	Tenant ConfigTenant
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
