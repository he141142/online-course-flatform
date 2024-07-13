package registry

import "sync"

type (
	Registrable interface {
		Key() string
	}

	Serializer interface {
		Serialize(v interface{}) ([]byte, error)
	}

	Deserializer interface {
		Deserialize(data []byte, v interface{}) error
	}

	Registry interface {
		Serialize(key string, v interface{}) ([]byte, error)
		Deserialize(key string, data []byte, v interface{}) error
		Build(key string) interface{}
		register(key string, fn func() interface{}, s Serializer, d Deserializer, opts ...BuildOption) error
	}
)

type registered struct {
	factory      func() interface{}
	serializer   Serializer
	deserializer Deserializer
}

type registry struct {
	registered map[string]registered
	mu         sync.RWMutex
}
