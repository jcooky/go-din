package din

type (
	Container struct {
		registry map[Name]any
		Env      Env
	}
	Env string
)

const (
	EnvProd Env = "prod"
	EnvTest Env = "test"
)

func NewContainer(env Env) *Container {
	return &Container{
		registry: map[Name]any{},
		Env:      env,
	}
}
