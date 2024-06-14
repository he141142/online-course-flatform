package ddd

type IDGetter interface {
	ID() string
}

type IDSetter interface {
	SetID(id string)
}

type Namer interface {
	GetEntityName() string
}

type NameSetter interface {
	SetName(name string)
}

type Entity interface {
	IDGetter
	IDSetter
	NameSetter
	Namer
}

type entity struct {
	id   string
	name string
}

var _ Entity = (*entity)(nil)

func NewEntity(id, name string) *entity {
	return &entity{
		name: name,
		id:   id,
	}
}

func (e *entity) GetEntityName() string {
	return e.name
}

func (e *entity) SetID(id string) {
	e.id = id
}

func (e *entity) ID() string {
	return e.id
}

func (e *entity) SetName(name string) {
	e.name = name
}