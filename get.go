package din

import (
	"fmt"
)

func Get[T any](c *Container, name Name) (res T, err error) {
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

	obj, err := fn(c)
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

func GetT[T any](c *Container) (res T, err error) {
	return Get[T](c, NewTypeName[T]())
}

func MustGet[T any](c *Container, name Name) T {
	res, err := Get[T](c, name)
	if err != nil {
		panic(fmt.Sprintf("error: %+v", err))
	}

	return res
}

func MustGetT[T any](c *Container) T {
	res, err := Get[T](c, NewTypeName[T]())
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

func SetT[T any](c *Container, obj T) {
	Set(c, NewTypeName[T](), obj)
}
