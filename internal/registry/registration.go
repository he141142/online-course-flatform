package registry

import "reflect"

func Register(r Registry, rg Registrable, s Serializer, d Deserializer, opts ...BuildOption) error {
	t := reflect.TypeOf(rg)
	var key string
	switch {
	case t.Kind() == reflect.Ptr && reflect.ValueOf(rg).IsNil():
		key = reflect.New(t).Interface().(Registrable).Key()
	default:
		key = rg.Key()
	}
	return RegisterKey(r, key, rg, s, d, opts...)
}

func RegisterKey(r Registry, key string, v interface{}, s Serializer, d Deserializer, opts ...BuildOption) error {
	t := reflect.TypeOf(v)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return r.register(key, func() interface{} {
		return reflect.New(t).Interface()
	}, s, d, opts...)
}
