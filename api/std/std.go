package std

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hellcats88/abstracte/api"
	"github.com/hellcats88/abstracte/logging"
	"github.com/hellcats88/abstracte/tenant"
)

type handlerCtx map[string]interface{}

type handler func(handlerCtx, http.ResponseWriter, *http.Request, *handlerInfo) bool

type handlerInfo struct {
	h    handler
	next *handlerInfo
}

func Next(ctx handlerCtx, w http.ResponseWriter, req *http.Request, info *handlerInfo) bool {
	if info != nil {
		return info.h(ctx, w, req, info.next)
	}

	return true
}

func array_interface2Handler(arr []interface{}) []handler {
	var result []handler
	for _, item := range arr {
		result = append(result, item.(handler))
	}

	return result
}

type group struct {
	name     string
	handlers []handler
}

type route struct {
	log      logging.Logger
	method   string
	handlers []handlerInfo
	service  api.Service
	ctx      handlerCtx
}

func (rt route) handle(w http.ResponseWriter, r *http.Request) {
	if rt.method != r.Method {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	rt.handlers[0].h(rt.ctx, w, r, rt.handlers[0].next)
}

type stdHttp struct {
	log       logging.Logger
	groups    map[string]group
	tenantCtx tenant.Context
	logCtx    logging.Context
}

func New(log logging.Logger) api.Http {
	return stdHttp{
		log:    log,
		groups: make(map[string]group),
	}
}

func (g *stdHttp) wrapService(service api.Service) {

}

func (g *stdHttp) defineBehavior(config api.Config) []handler {
	var handlers []handler
	handlers = append(handlers, g.createLogContextFromHeaders)

	switch config.Tenant {
	case api.ConfigTenantFromHeaders:
		handlers = append(handlers, g.retrieveTenant)
	}

	return handlers
}

func (g *stdHttp) retrieveTenant(ctx handlerCtx, w http.ResponseWriter, r *http.Request, info *handlerInfo) bool {
	tCtx := tenant.Context{
		Id:     r.Header.Get("X-Tenant-ID"),
		UserId: r.Header.Get("X-Tenant-UserID"),
	}

	if tCtx.Id == "" || tCtx.UserId == "" {
		g.log.Error(g.logCtx, "Rejected request caused by missing tenant informations")

		bt, _ := json.Marshal(api.Model{
			Error: api.ErrorModel{
				Code:   api.ApiErrorAuthFailed,
				Msg:    "Failed to get user information",
				DevMsg: "Missing X-Tenant-ID or X-Tenant-UserID headers",
			},
		})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		if _, err := w.Write(bt); err != nil {
			g.log.Error(g.logCtx, "Failed to write response not authorized to client. %v", err)
		}

		return false
	}

	g.tenantCtx = tCtx

	ctx[api.TenantKey] = tCtx
	return Next(ctx, w, r, info)
}

func (g *stdHttp) createLogContextFromHeaders(ctx handlerCtx, w http.ResponseWriter, r *http.Request, info *handlerInfo) bool {
	var lCtx logging.Context

	if corrId := r.Header.Get("X-Correlation-ID"); corrId != "" {
		lCtx = logging.NewContext(corrId)
	} else {
		lCtx = logging.NewContextUUID()
	}

	g.logCtx = lCtx

	ctx[api.LogKey] = lCtx
	return Next(ctx, w, r, info)
}

func (g stdHttp) AddGroup(name string, subPath string, config api.ConfigGroup) error {
	if _, exists := g.groups[name]; exists {
		return fmt.Errorf("Group %s already exists, skipping creation", name)
	}

	var handlers []handler
	handlers = append(handlers, g.defineBehavior(config.Config)...)
	handlers = append(handlers, array_interface2Handler(config.Handlers)...)

	g.groups[name] = group{
		name:     name,
		handlers: handlers,
	}

	return nil
}

func (g stdHttp) AddGroupRoute(method string, path string, group string, config api.ConfigRoute, service api.Service) error {
	groupRef, exists := g.groups[group]
	if !exists {
		return fmt.Errorf("Group %s does not exists, skipping route %s creation", group, path)
	}

	http.HandleFunc(path, func(rw http.ResponseWriter, r *http.Request) {
		var handlers []handler
		handlers = append(handlers, groupRef.handlers...)
		handlers = append(handlers, g.defineBehavior(config.Config)...)
		handlers = append(handlers, array_interface2Handler(config.CustomHandlers)...)

		var hInfo []handlerInfo

		for _, h := range handlers {
			hInfo = append(hInfo, handlerInfo{
				h: h,
			})
		}

		//last handler have next = nil so we can cycle to len - 1
		for i := 0; i < len(hInfo)-1; i++ {
			hInfo[i].next = &hInfo[i+1]
		}

		route{
			log:      g.log,
			method:   method,
			handlers: hInfo,
			service:  service,
			ctx:      make(handlerCtx),
		}.handle(rw, r)
	})

	return nil
}

func (g stdHttp) AddRoute(method string, path string, config api.ConfigRoute, service api.Service) error {
	var handlers []handler
	handlers = append(handlers, g.defineBehavior(config.Config)...)
	handlers = append(handlers, array_interface2Handler(config.CustomHandlers)...)

	var hInfo []handlerInfo

	for _, h := range handlers {
		hInfo = append(hInfo, handlerInfo{
			h: h,
		})
	}

	//last handler have next = nil so we can cycle to len - 1
	for i := 0; i < len(hInfo)-1; i++ {
		hInfo[i].next = &hInfo[i+1]
	}

	http.HandleFunc(path, func(rw http.ResponseWriter, r *http.Request) {
		route{
			log:      g.log,
			method:   method,
			handlers: hInfo,
			service:  service,
			ctx:      make(handlerCtx),
		}.handle(rw, r)
	})

	return nil
}

func (g stdHttp) Listen(port int, address string) error {
	return http.ListenAndServe(fmt.Sprintf("%s:%d", address, port), nil)
}
