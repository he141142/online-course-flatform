package ddd

const (
	AggNameKey    = "aggregate.name"
	AggIDKey      = "aggregate.id"
	AggVersionKey = "aggregate.version"
)

type (
	aggregate struct {
		Entity
		events []AggregateEvent
	}

	AggregateNamer interface {
		GetAggregateName() string
	}

	// Eventer enable domain with events
	Eventer interface {
		AddEvent(string, EventPayload, ...EventOption)
		Events() []AggregateEvent
		ClearEvents()
	}

	aggregateEvent struct {
		event
	}

	AggregateEvent interface {
		Event
		AggregateName() string
		AggregateID() string
		AggregateVersion() string
	}

	Aggregate interface {
		IDGetter
		AggregateNamer
		Eventer
		IDSetter
		NameSetter
	}
)

var (
	_ AggregateEvent = (*aggregateEvent)(nil)
	_ Aggregate      = (*aggregate)(nil)
)

func (a *aggregate) GetAggregateName() string {
	return a.GetEntityName()
}

func (a *aggregate) AddEvent(eventName string, payload EventPayload, options ...EventOption) {
	options = append(options, Metadata{
		AggIDKey:   a.ID(),
		AggNameKey: a.GetEntityName(),
	})

	a.events = append(a.events,
		&aggregateEvent{
			event: *newEvent(eventName, payload, options...),
		})
}

func (a *aggregate) Events() []AggregateEvent {
	return a.events
}

func (a *aggregate) ClearEvents() {
	a.events = make([]AggregateEvent, 0)
}

func (ae *aggregateEvent) AggregateID() string {
	return ae.metadata.ValueByKey(AggIDKey).(string)
}

func (ae *aggregateEvent) AggregateName() string {
	return ae.metadata.ValueByKey(AggNameKey).(string)
}

func (ae *aggregateEvent) AggregateVersion() string {
	return ae.metadata.ValueByKey(AggVersionKey).(string)
}