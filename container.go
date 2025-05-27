package din

import (
	"context"
	"slices"
)

type (
	ShutdownFunc func(ctx context.Context)
	Container    struct {
		context.Context
		registry      map[Name]any
		Env           Env
		shutdownFuncs []ShutdownFunc
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

func (c *Container) RegisterOnShutdown(fn ShutdownFunc) {
	c.shutdownFuncs = append(c.shutdownFuncs, fn)
}

func (c *Container) Close() {
	ctx := context.WithoutCancel(c.Context)
	for _, fn := range slices.Backward(c.shutdownFuncs) {
		fn(ctx)
	}
}
