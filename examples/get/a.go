package get

import (
	"github.com/jcooky/go-din"
)

type A interface {
	Foo() string
}

type a struct {
	foo string
}

func (a *a) Foo() string {
	return a.foo
}

func init() {
	din.RegisterT(func(c *din.Container) (A, error) {
		return &a{
			foo: "foo",
		}, nil
	})
}
