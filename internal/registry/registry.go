package registry

import (
	"fmt"
	"sync"
)

type (
	ErrKeyAlreadyExisted string
	ErrKeyNotExisted     string
)

func (s ErrKeyAlreadyExisted) Error() string {
	return fmt.Sprintf("err key %s alreadyExisted", string(s))
}

func (s ErrKeyNotExisted) Error() string {
	return fmt.Sprintf("err key %s does not existed", string(s))
}

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
		Deserialize(key string, data []byte, opts ...BuildOption) (interface{}, error)
		Build(key string, opts ...BuildOption) (interface{}, error)
		register(
			key string,
			fn func() interface{},
			s Serializer,
			d Deserializer,
			opts ...BuildOption,
		) error
	}
)

type registered struct {
	factory      func() interface{}
	serializer   Serializer
	deserializer Deserializer
}

var _ Registry = (*registry)(nil)

type registry struct {
	registered map[string]registered
	mu         sync.RWMutex
}

func New() Registry {
	return &registry{
		registered: make(map[string]registered),
		mu:         sync.RWMutex{},
	}
}

func (r *registry) register(
	key string,
	fn func() interface{},
	s Serializer,
	d Deserializer,
	opts ...BuildOption,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, existed := r.registered[key]; existed {
		return ErrKeyAlreadyExisted(key)
	}

	r.registered[key] = registered{
		factory:      fn,
		serializer:   s,
		deserializer: d,
	}
	return nil
}

func (r *registry) Build(key string, opts ...BuildOption) (interface{}, error) {
	if _, existed := r.registered[key]; !existed {
		return nil, ErrKeyNotExisted(key)
	}
	v := r.registered[key].factory()

	for _, opt := range opts {
		err := opt(v)
		if err != nil {
			return nil, err
		}
	}

	return v, nil
}

func (r *registry) Serialize(key string, v interface{}) ([]byte, error) {
	if !r.DoesKeyExisted(key) {
		return nil, ErrKeyNotExisted(key)
	}
	return r.registered[key].serializer.Serialize(v)
}

func (r *registry) DoesKeyExisted(key string) bool {
	_, existed := r.registered[key]

	return existed
}

func (r *registry) Deserialize(key string, data []byte, opts ...BuildOption) (interface{}, error) {
	v, err := r.Build(key, opts...)
	if err != nil {
		return nil, err
	}
	if !r.DoesKeyExisted(key) {
		return nil, ErrKeyNotExisted(key)
	}
	err = r.registered[key].deserializer.Deserialize(data, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
