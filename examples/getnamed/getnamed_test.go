package getnamed_test

import (
	"context"
	"github.com/jcooky/go-din"
	"github.com/jcooky/go-din/examples/getnamed"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestB_Bar(t *testing.T) {
	ctx := context.TODO()
	container := din.NewContainer(ctx, din.EnvTest)

	b, err := din.Get[*getnamed.B](container, getnamed.BName)
	require.Nil(t, err)

	value := b.Bar()
	require.Equal(t, "foo and bar", value)
}
