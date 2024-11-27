package design_pattern

type Event interface {
	Type() string
	Data() interface{}
}

type BaseEvent struct {
	EventType string
	EventData interface{}
}

func (e *BaseEvent) Type() string {
	return e.EventType
}

func (e *BaseEvent) Data() interface{} {
	return e.EventData
}
