package getnamed

import (
	"github.com/jcooky/go-din"
)

type A interface {
	Foo() string
}

type a struct {
	foo string
}

var (
	AName   = din.NewRandomName()
	_     A = (*a)(nil)
)

func (a *a) Foo() string {
	return a.foo
}

func init() {
	din.Register(AName, func(c *din.Container) (any, error) {
		return &a{
			foo: "foo",
		}, nil
	})
}
