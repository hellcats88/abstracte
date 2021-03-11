package api

type ConfigLog uint

const (
	ConfigLogRandom      ConfigLog = 0x0
	ConfigLogFromHeaders ConfigLog = 0x1
)

type ConfigTenant uint

const (
	ConfigTenantNo          ConfigTenant = 0x0
	ConfigTenantFromHeaders ConfigTenant = 0x1
)

type ApiConfig struct {
	CustomPreHandlers  []interface{}
	CustomMidHandlers  []interface{}
	CustomPostHandlers []interface{}
	Tenant             ConfigTenant
}
