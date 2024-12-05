package wraps

import (
	"net/http"
	"reflect"

	"github.com/ogiusek/hw/src/hw"
)

type Injectable interface {
	FromHttp(*http.Request) (any, hw.Resp)
}

func Inject[T any](fn hw.Wrapper[T]) hw.Wrapper[T] {
	return func(args T, r *http.Request) any {
		// argsReflection := argsPtrReflection.Elem()
		argsReflection := reflect.ValueOf(&args).Elem()
		for i := 0; i < argsReflection.NumField(); i++ {
			field := argsReflection.Field(i)
			if injectable, ok := field.Interface().(Injectable); ok {
				value, resp := injectable.FromHttp(r)
				if resp != nil {
					return resp
				}
				field.Set(reflect.ValueOf(value))
			}
		}
		res := fn(args, r)

		return res
	}
}
