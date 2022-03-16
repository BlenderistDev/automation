package interfaces

type TriggerEvent interface {
	GetName() string
	GetData() map[string]string
}

type Condition interface {
	Check(trigger TriggerEvent) (bool, error)
}

type Action interface {
	Execute(trigger TriggerEvent) error
}

type Automation interface {
	Execute(trigger TriggerEvent) error
	AddTrigger(trigger string)
	GetTriggers() []string
	AddAction(action Action)
	AddCondition(condition Condition)
}
