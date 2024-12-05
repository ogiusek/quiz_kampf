package wraps

import (
	"lib/common/errs"
	"net/http"

	"github.com/ogiusek/hw/src/hw"
)

func Catch[T any](fn hw.Wrapper[T]) hw.Wrapper[T] {
	return func(args T, r *http.Request) any {
		res := fn(args, r)
		if err, ok := res.(error); ok {
			w := hw.NewResponse()
			errs.ToHttp(w, err)
			return w
		}
		return res
	}
}
