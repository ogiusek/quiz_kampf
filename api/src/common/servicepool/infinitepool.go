package servicepool

type infiniteServicePool[T any] struct{ get func() T }

func (services *infiniteServicePool[T]) Get() (T, func()) {
	return services.get(), func() {}
}

func NewInfinitePool[T any](fn func() T) ServicePool[T] {
	return &infiniteServicePool[T]{get: fn}
}
