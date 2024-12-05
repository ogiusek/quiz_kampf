package wraps

import (
	"net/http"

	"github.com/ogiusek/hw/src/hw"
)

type iResponse interface {
	Response() hw.Resp
}

func Response[T any](fn hw.Wrapper[T]) hw.Wrapper[T] {
	return func(args T, r *http.Request) any {
		res := fn(args, r)
		if resp, ok := res.(iResponse); ok {
			return resp.Response()
		}
		return res
	}
}
