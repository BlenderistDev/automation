package interfaces

// TriggerEvent interface for automation trigger events
type TriggerEvent interface {
	GetName() string
	GetData() map[string]string
}

// Condition automation condition interface
type Condition interface {
	Check(trigger TriggerEvent) (bool, error)
}

// Action automation condition interface
type Action interface {
	Execute(trigger TriggerEvent) error
}

// Automation automation interface
type Automation interface {
	Execute(trigger TriggerEvent) error
	AddTrigger(trigger string)
	GetTriggers() []string
	AddAction(action Action)
	AddCondition(condition Condition)
}
