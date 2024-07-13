package serdes

import "drake.elearn-platform.ru/internal/registry"

type jsonSerde struct {
	r registry.Registry
}

func (j jsonSerde) Register(registrable registry.Registrable, serializer registry.Serializer, deserializer registry.Deserializer, option ...registry.BuildOption) error {
	return r.re
}

func (j jsonSerde) RegisterKey(s string, s2 string, option ...registry.BuildOption) error {
	//TODO implement me
	panic("implement me")
}

func NewJSONSerde(r registry.Registry) registry.Serde {
	return &jsonSerde{r: r}
}
