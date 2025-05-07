package din

import (
	"context"
	"fmt"
)

func Get[T any](ctx context.Context, c *Container, name Name) (res T, err error) {
	if c == nil {
		err = fmt.Errorf("container should be not nil")
		return
	}

	var ok bool
	res, ok = c.registry[name].(T)
	if ok {
		return
	}

	fn, ok := g[name]
	if !ok {
		err = fmt.Errorf("object %s not registered", name)
		return
	}

	obj, err := fn(ctx, c)
	if err != nil {
		return
	}

	c.registry[name] = obj

	res, ok = obj.(T)
	if !ok {
		err = fmt.Errorf("object %s is not of type %T", name, res)
		return
	}

	return
}

func GetT[T any](ctx context.Context, c *Container) (res T, err error) {
	return Get[T](ctx, c, NewTypeName[T]())
}

func MustGet[T any](ctx context.Context, c *Container, name Name) T {
	res, err := Get[T](ctx, c, name)
	if err != nil {
		panic(fmt.Sprintf("error: %+v", err))
	}

	return res
}

func MustGetT[T any](ctx context.Context, c *Container) T {
	res, err := Get[T](ctx, c, NewTypeName[T]())
	if err != nil {
		panic(fmt.Sprintf("error: %+v", err))
	}

	return res
}

func Set(c *Container, name Name, obj any) {
	if c == nil {
		panic("container should be not nil")
	}

	c.registry[name] = obj
}

func SetT[T any](c *Container, obj any) {
	Set(c, NewTypeName[T](), obj)
}
