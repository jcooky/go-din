package din

import "context"

type (
	RegisterTFn[T any] func(context.Context, *Container) (T, error)
	RegisterFn         = RegisterTFn[any]
)

var (
	g = make(map[Name]RegisterFn)
)

func Register(name Name, fn RegisterFn) {
	g[name] = fn
}

func RegisterT[T any](fn RegisterTFn[T]) {
	g[NewTypeName[T]()] = func(ctx context.Context, c *Container) (any, error) {
		return fn(ctx, c)
	}
}
