package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hellcats88/abstracte/api"
	"github.com/hellcats88/abstracte/logging"
	"github.com/hellcats88/abstracte/tenant"
)

func array_interface2HandlerFunc(arr []interface{}) []gin.HandlerFunc {
	var result []gin.HandlerFunc
	for _, item := range arr {
		result = append(result, item.(gin.HandlerFunc))
	}

	return result
}

type ginHttp struct {
	engine    *gin.Engine
	log       logging.Logger
	groups    map[string]*gin.RouterGroup
	tenantCtx tenant.Context
	logCtx    logging.Context
}

func New(log logging.Logger) api.Http {
	engine := gin.Default()
	return ginHttp{
		engine: engine,
		log:    log,
		groups: make(map[string]*gin.RouterGroup),
	}
}

func (g *ginHttp) wrapService(service api.Service) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (g *ginHttp) definePreBehavior(config api.Config) []gin.HandlerFunc {
	var handlers []gin.HandlerFunc
	handlers = append(handlers, func(c *gin.Context) { g.createLogContextFromHeaders(c) })

	switch config.Tenant {
	case api.ConfigTenantFromHeaders:
		handlers = append(handlers, func(c *gin.Context) { g.retrieveTenant(c) })
	}

	return handlers
}

func (g *ginHttp) definePostBehavior(config api.Config) []gin.HandlerFunc {
	var handlers []gin.HandlerFunc
	return handlers
}

func (g *ginHttp) retrieveTenant(ctx *gin.Context) {
	tCtx := tenant.Context{
		Id:     ctx.GetHeader("X-Tenant-ID"),
		UserId: ctx.GetHeader("X-Tenant-UserID"),
	}

	if tCtx.Id == "" || tCtx.UserId == "" {
		g.log.Error(g.logCtx, "Rejected request caused by missing tenant informations")

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.Model{
			Error: api.ErrorModel{
				Code:   api.ApiErrorAuthFailed,
				Msg:    "Failed to get user information",
				DevMsg: "Missing X-Tenant-ID or X-Tenant-UserID headers",
			},
		})
	}

	g.tenantCtx = tCtx

	ctx.Set(api.TenantKey, tCtx)
	ctx.Next()
}

func (g *ginHttp) createLogContextFromHeaders(ctx *gin.Context) {
	var lCtx logging.Context

	if corrId := ctx.GetHeader("X-Correlation-ID"); corrId != "" {
		lCtx = logging.NewContext(corrId)
	} else {
		lCtx = logging.NewContextUUID()
	}

	g.logCtx = lCtx

	ctx.Set(api.LogKey, lCtx)
	ctx.Next()
}

func (g ginHttp) AddGroup(name string, subPath string, config api.ConfigGroup) error {
	if _, exists := g.groups[name]; exists {
		return fmt.Errorf("Group %s already exists, skipping creation", name)
	}

	var handlers []gin.HandlerFunc
	handlers = append(handlers, g.definePreBehavior(config.Config)...)
	handlers = append(handlers, array_interface2HandlerFunc(config.CustomHandlers)...)
	handlers = append(handlers, g.definePostBehavior(config.Config)...)

	g.groups[name] = g.engine.Group(subPath, handlers...)

	return nil
}

func (g ginHttp) AddGroupRoute(method string, path string, group string, config api.ConfigRoute, service api.Service) error {
	groupRef, exists := g.groups[group]
	if !exists {
		return fmt.Errorf("Group %s does not exists, skipping route %s creation", group, path)
	}

	var handlers []gin.HandlerFunc
	handlers = append(handlers, g.definePreBehavior(config.Config)...)
	handlers = append(handlers, g.wrapService(service))
	handlers = append(handlers, g.definePostBehavior(config.Config)...)

	groupRef.Handle(method, path, handlers...)
	return nil
}

func (g ginHttp) AddRoute(method string, path string, config api.ConfigRoute, service api.Service) error {
	var handlers []gin.HandlerFunc
	handlers = append(handlers, g.definePreBehavior(config.Config)...)
	handlers = append(handlers, g.wrapService(service))
	handlers = append(handlers, g.definePostBehavior(config.Config)...)

	g.engine.Handle(method, path, handlers...)
	return nil
}

func (g ginHttp) Listen(port int, address string) error {
	return g.engine.Run(fmt.Sprintf("%s:%d", address, port))
}
