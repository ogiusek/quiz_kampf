package httpapi

import (
	"errors"
	"net/http"
)

type EndpointFunc func(w http.ResponseWriter, r *http.Request)
type MiddlewareFunc func(w http.ResponseWriter, r *http.Request, next func())

type api struct {
	// data
	endpointsPrefix UrlPrefix
	endpoints       map[Url]map[Method]EndpointFunc
	middlewares     []MiddlewareFunc
	// methods
	router Router // url matching
}

type IApi interface {
	Map(url Url, method Method, fn EndpointFunc) error
	Use(middleware MiddlewareFunc)
	Prefix(prefix UrlPrefix)
	Router(router Router)
}

func New() api {
	return api{
		endpointsPrefix: "",
		endpoints:       map[Url]map[Method]EndpointFunc{},
		middlewares:     []MiddlewareFunc{},
		router:          defaultRouter,
	}
}

func (api *api) Map(url Url, method Method, fn EndpointFunc) error {
	if api.endpoints[url][method] != nil {
		return errors.New("endpoint is already used")
	}

	if api.endpoints[url] == nil {
		api.endpoints[url] = map[Method]EndpointFunc{}
	}

	api.endpoints[url][method] = fn
	return nil
}

func (api *api) Use(middleware MiddlewareFunc) {
	api.middlewares = append(api.middlewares, middleware)
}

func (api *api) Prefix(prefix UrlPrefix) {
	api.endpointsPrefix = prefix
}

func (api *api) Router(router Router) {
	api.router = router
}
