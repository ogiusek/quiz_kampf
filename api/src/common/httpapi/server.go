package httpapi

import (
	"net/http"
	"strings"
)

func runEndpoint(w http.ResponseWriter, r *http.Request, middlewares []MiddlewareFunc, endpoint *EndpointFunc) {
	if len(middlewares) == 0 {
		(*endpoint)(w, r)
		return
	}
	middleware := middlewares[0]
	middleware(w, r, func() {
		runEndpoint(w, r, middlewares[1:], endpoint)
	})
}

func (api api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestPath := RequestPath(r.RequestURI)
	paramsIndex := strings.Index(string(requestPath), "?")
	if paramsIndex != -1 {
		requestPath = requestPath[:paramsIndex]
	}
	requestMethod := Method(r.Method)

	for path, methods := range api.endpoints {
		if api.router(api.endpointsPrefix, path, requestPath) {
			for method, endpoint := range methods {
				if method == requestMethod {
					runEndpoint(w, r, api.middlewares, &endpoint)
					return
				}
			}
		}
	}

	w.WriteHeader(404)
}
