package wraps

import (
	"lib/common/errs"
	"lib/common/valid"
	"net/http"

	"github.com/ogiusek/hw/src/hw"
)

func Validate[T any](fn hw.Wrapper[T]) hw.Wrapper[T] {
	return func(args T, r *http.Request) any {
		if err := valid.Valid(args); err != nil {
			w := hw.NewResponse()
			errs.ToHttp(w, errs.InvalidInput(err.Error()))
			return w
		}
		return fn(args, r)
	}
}
