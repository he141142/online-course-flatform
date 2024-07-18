package serdes

import (
	"bytes"
	"encoding/json"

	"drake.elearn-platform.ru/internal/registry"
)

type jsonSerde struct {
	r registry.Registry
}

type (
	JsonDeserializer func(data []byte, v interface{}) error
	JsonSerializer   func(v interface{}) ([]byte, error)
)

func (deserializer JsonDeserializer) Deserialize(data []byte, v interface{}) error {
	return deserializer(data, v)
}
func (serializer JsonSerializer) Serialize(v interface{}) ([]byte, error) {
	return serializer(v)
}

func newDeserializer() JsonDeserializer {
	return func(data []byte, v interface{}) error {
		b := bytes.NewBuffer(data)
		return json.NewDecoder(b).Decode(v)
	}
}

func newSerializer() JsonSerializer {
	return func(v interface{}) ([]byte, error) {
		//var buff io.Writer
		//var out []byte
		//err := json.NewEncoder(buff).Encode(out)
		//if err != nil {
		//	return nil, err
		//}
		//return out, nil
		return json.Marshal(v)
	}
}
func (j jsonSerde) Register(
	registrable registry.Registrable,
	option ...registry.BuildOption,
) error {
	return registry.Register(j.r, registrable, j.buildSerializer(), j.buildDeserializer(), option...)
}

func (j jsonSerde) RegisterKey(
	key string,
	v interface{},
	option ...registry.BuildOption,
) error {
	return registry.RegisterKey(j.r, key, v, j.buildSerializer(), j.buildDeserializer(), option...)
}

func NewJSONSerde(r registry.Registry) registry.Serde {
	return &jsonSerde{r: r}
}

func (j jsonSerde) buildDeserializer() registry.Deserializer {
	return newDeserializer()
}

func (j jsonSerde) buildSerializer() registry.Serializer {
	return newSerializer()
}
