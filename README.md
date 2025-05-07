# go-din [![Go Test & Lint](https://github.com/jcooky/go-din/actions/workflows/ci.yml/badge.svg)](https://github.com/jcooky/go-din/actions/workflows/ci.yml)

A lightweight dependency injection container for Go applications, leveraging Go generics for type-safe dependency management.

## Features

- Type-safe dependency injection using Go generics
- Simple API for registering, retrieving, and managing dependencies
- Support for different environments (production, testing)
- Easy mocking of dependencies for testing
- Both named and type-based dependency resolution

## Installation

```bash
go get -u github.com/jcooky/go-din
```

## How to use

### Creating a Container

```go
// Create a new container with production environment
container := din.NewContainer(din.EnvProd)

// Or for testing
container := din.NewContainer(din.EnvTest)
```

### Registering Dependencies

You can register dependencies in two ways:

1. By type (recommended):

```go
// Register a dependency factory for type A
din.RegisterT(func(ctx context.Context, c *din.Container) (A, error) {
    return &a{
        foo: "foo",
    }, nil
})
```

2. By name:

```go
// Register a dependency with a custom name
din.Register(din.Name("myService"), func(ctx context.Context, c *din.Container) (any, error) {
    return &MyService{}, nil
})
```

### Retrieving Dependencies

Retrieve dependencies type-safely:

```go
// Get by type (recommended)
a, err := din.GetT[A](ctx, container)
if err != nil {
    // Handle error
}

// Or with panic on error
a := din.MustGetT[A](ctx, container)

// Get by name
a, err := din.Get[A](ctx, container, din.Name("myService"))
if err != nil {
    // Handle error
}
```

### Managing Dependencies with Dependencies

Dependencies can depend on other dependencies:

```go
din.RegisterT(func(ctx context.Context, c *din.Container) (*B, error) {
    // Get dependency A
    a := din.MustGetT[A](ctx, c)
    
    // Create B with A as a dependency
    b := &B{
        bar: "bar",
        a:   a,
    }

    return b, nil
})
```

### Testing with Mocks

Easily replace dependencies for testing:

```go
// Create test container
container := din.NewContainer(din.EnvTest)

// Set mock implementation
din.SetT[ServiceInterface](container, &MockService{})

// Get component with mocked dependency
component, err := din.GetT[*Component](ctx, container)
```

## Example

```go
// Define interfaces and implementations
type A interface {
    Foo() string
}

type a struct {
    foo string
}

func (a *a) Foo() string {
    return a.foo
}

// Register implementation
func init() {
    din.RegisterT(func(ctx context.Context, c *din.Container) (A, error) {
        return &a{foo: "foo"}, nil
    })
}

// Component that depends on A
type B struct {
    bar string
    a   A
}

func (b *B) Bar() string {
    return fmt.Sprintf("%s and %s", b.a.Foo(), b.bar)
}

// Register component with dependency
func init() {
    din.RegisterT(func(ctx context.Context, c *din.Container) (*B, error) {
        a := din.MustGetT[A](ctx, c)
        b := &B{
            bar: "bar",
            a:   a,
        }
        return b, nil
    })
}

// Usage
func main() {
    ctx := context.Background()
    container := din.NewContainer(din.EnvProd)
    
    b, err := din.GetT[*B](ctx, container)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(b.Bar()) // Output: foo and bar
}
```

See more detailed examples in the [examples](examples) directory.

## License

[MIT License](LICENSE)