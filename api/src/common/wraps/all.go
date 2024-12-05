package wraps

import (
	"net/http"

	"github.com/ogiusek/hw/src/hw"
)

func Wrap[T any](fn hw.Wrapper[T]) func(http.ResponseWriter, *http.Request) {
	return hw.Wrap(Catch(Inject(Validate(Response(fn)))), nil)
}
