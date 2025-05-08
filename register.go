package din

type (
	RegisterTFn[T any] func(*Container) (T, error)
	RegisterFn         = RegisterTFn[any]
)

var (
	g = make(map[Name]RegisterFn)
)

func Register(name Name, fn RegisterFn) {
	g[name] = fn
}

func RegisterT[T any](fn RegisterTFn[T]) {
	g[NewTypeName[T]()] = func(c *Container) (any, error) {
		return fn(c)
	}
}
