package factorymethod

// 辅助定义
type Type uint8

const (
	Start Type = iota
	End
)

type Event interface {
	EventType() Type
}

type StartEvent struct{}

type EndEvent struct{}

func (s *StartEvent) EventType() Type {
	return Start
}

func (e *EndEvent) EventType() Type {
	return End
}

// 工厂对象实现
type Factory struct{}

func (f *Factory) Create(eType Type) Event {
	switch eType {
	case Start:
		return &StartEvent{}
	case End:
		return &EndEvent{}
	default:
		return nil
	}
}

// 工厂方法实现
func OfStart() Event {
	return &StartEvent{}
}

func OfEnd() Event {
	return &EndEvent{}
}
