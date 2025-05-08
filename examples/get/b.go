package get

import (
	"fmt"
	"github.com/jcooky/go-din"
)

type B struct {
	bar string
	a   A
}

func (b *B) Bar() string {
	return fmt.Sprintf("%s and %s", b.a.Foo(), b.bar)
}

func init() {
	din.RegisterT(func(c *din.Container) (*B, error) {
		a := din.MustGetT[A](c)
		b := &B{
			bar: "bar",
			a:   a,
		}

		return b, nil
	})
}
