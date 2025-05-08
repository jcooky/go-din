package din_test

import (
	"context"
	"github.com/jcooky/go-din"
	"testing"
)

type (
	A interface {
		Foo() string
	}

	a struct {
		foo string
	}
)

func (a *a) Foo() string {
	return a.foo
}

func TestSetter(t *testing.T) {
	ctx := context.TODO()
	c := din.NewContainer(ctx, din.EnvTest)

	din.SetT[A](c, &a{})
}
