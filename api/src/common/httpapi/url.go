package httpapi

import (
	"strings"
)

type RequestPath string
type Url string
type UrlPrefix string
type Method string

type Router func(prefix UrlPrefix, url Url, requestPath RequestPath) bool

func defaultRouter(prefix UrlPrefix, url Url, requestPath RequestPath) bool {
	path := ""

	path += strings.Trim(string(prefix), "/")
	path += "/"
	path += strings.Trim(string(url), "/")

	path = strings.Trim(path, "/")
	expected := strings.Trim(string(requestPath), "/")

	return path == expected
}
