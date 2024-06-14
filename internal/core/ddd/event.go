package ddd

import (
	"time"

	"github.com/google/uuid"
)

type (
	event struct {
		Entity
		payload   EventPayload
		metadata  Metadata
		ocurredAt time.Time
	}

	// EventOption, example: use Metadata to configure event
	EventOption interface {
		configureEvent(e *event)
	}

	Event interface {
		IDGetter
		EventName() string
		Payload() EventPayload
		Metadata() Metadata
		OcurredAt() time.Time
	}

	EventPayload interface{}
)

var _ Event = (*event)(nil)

func NewEvent(name string, payload EventPayload, options ...EventOption) Event {
	return newEvent(name, payload, options...)
}

// newEvent for internal usage, return struct unpublished
func newEvent(name string, payload EventPayload, options ...EventOption) *event {
	evt := &event{
		Entity:    NewEntity(uuid.New().String(), name),
		payload:   payload,
		metadata:  make(Metadata),
		ocurredAt: time.Now(),
	}
	for _, opt := range options {
		opt.configureEvent(evt)
	}
	return evt
}

func (e *event) ID() string {
	return e.Entity.ID()
}

func (e *event) EventName() string {
	return e.GetEntityName()
}

func (e *event) Payload() EventPayload {
	return e.payload
}

func (e *event) Metadata() Metadata {
	return e.metadata
}

func (e *event) OcurredAt() time.Time {
	return e.ocurredAt
}