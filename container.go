package din

import (
	"context"
	"slices"
)

type (
	ShutdownFunc func(ctx context.Context)
	Container    interface {
		WithContext(ctx context.Context) Container
		Get(name Name) any
		Set(name Name, obj any)
		Close()
		RegisterOnShutdown(fn ShutdownFunc)
		Env() Env
	}
	container struct {
		context.Context
		registry      map[Name]any
		env           Env
		shutdownFuncs []ShutdownFunc
	}
	Env string
)

const (
	EnvProd Env = "prod"
	EnvTest Env = "test"
)

func NewContainer(baseCtx context.Context, env Env) Container {
	if baseCtx == nil {
		baseCtx = context.Background()
	}
	return &container{
		Context:  baseCtx,
		registry: map[Name]any{},
		env:      env,
	}
}

func (c *container) Env() Env {
	return c.env
}

func (c *container) Get(name Name) any {
	return c.registry[name]
}

func (c *container) Set(name Name, obj any) {
	c.registry[name] = obj
}

func (c *container) WithContext(ctx context.Context) Container {
	if ctx == nil {
		ctx = context.Background()
	}
	return &container{
		Context:  ctx,
		registry: c.registry,
		env:      c.env,
	}
}

func (c *container) RegisterOnShutdown(fn ShutdownFunc) {
	c.shutdownFuncs = append(c.shutdownFuncs, fn)
}

func (c *container) Close() {
	ctx := context.WithoutCancel(c.Context)
	for _, fn := range slices.Backward(c.shutdownFuncs) {
		fn(ctx)
	}
}
