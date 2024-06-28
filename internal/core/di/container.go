package di

type Scope int

const (
	// Singleton is a scope that creates a single instance of the type.
	Singleton Scope = iota + 1
	Scoped
)

type contextKey int

const containerKey contextKey = iota + 1

type Container interface {
}

type DepFactoryFunc func(container Container) (any, error)
