package din

import (
	"context"
)

type (
	Container struct {
		context.Context
		registry map[Name]any
		Env      Env
	}
	Env string
)

const (
	EnvProd Env = "prod"
	EnvTest Env = "test"
)

func NewContainer(baseCtx context.Context, env Env) *Container {
	if baseCtx == nil {
		baseCtx = context.Background()
	}
	return &Container{
		Context:  baseCtx,
		registry: map[Name]any{},
		Env:      env,
	}
}

func (c *Container) WithContext(ctx context.Context) *Container {
	if ctx == nil {
		ctx = context.Background()
	}
	return &Container{
		Context:  ctx,
		registry: c.registry,
		Env:      c.Env,
	}
}
