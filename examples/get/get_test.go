package get_test

import (
	"context"
	"github.com/jcooky/go-din"
	"github.com/jcooky/go-din/examples/get"
	"github.com/stretchr/testify/require"
	"testing"
)

type (
	AMock struct{}
)

func (a *AMock) Foo() string {
	return "mocked foo"
}

func TestB_Bar(t *testing.T) {
	ctx := context.TODO()
	container := din.NewContainer(ctx, din.EnvTest)

	b, err := din.GetT[*get.B](container)
	require.Nil(t, err)

	value := b.Bar()
	require.Equal(t, "foo and bar", value)
}

func TestBWithAMock(t *testing.T) {
	ctx := context.TODO()
	container := din.NewContainer(ctx, din.EnvTest)

	din.SetT[get.A](container, &AMock{})
	b, err := din.GetT[*get.B](container)
	require.Nil(t, err)

	value := b.Bar()
	require.Equal(t, "mocked foo and bar", value)
}
