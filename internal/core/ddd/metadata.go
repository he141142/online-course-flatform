package ddd

type Metadata map[string]any

func (m Metadata) ValueByKey(key string) any {
	return m[key]
}

func (m Metadata) configureEvent(e *event) {
	for k, v := range m {
		e.metadata[k] = v
	}
}