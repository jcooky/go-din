package getnamed

import (
	"fmt"
	"github.com/jcooky/go-din"
)

type B struct {
	bar string
	a   A
}

var (
	BName = din.NewRandomName()
)

func (b *B) Bar() string {
	return fmt.Sprintf("%s and %s", b.a.Foo(), b.bar)
}

func init() {
	din.Register(BName, func(c *din.Container) (any, error) {
		a := din.MustGet[A](c, AName)
		b := &B{
			bar: "bar",
			a:   a,
		}

		return b, nil
	})
}
