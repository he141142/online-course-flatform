package registry

type Serde interface {
	Register(Registrable, Serializer, Deserializer, ...BuildOption) error
	RegisterKey(string, string, ...BuildOption) error
}
