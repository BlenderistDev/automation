package interfaces

// TriggerEvent interface for automation trigger events
type TriggerEvent interface {
	// GetName returns trigger event name
	GetName() string
	// GetData returns trigger event data
	GetData() map[string]string
	// GetFromMeta return value from trigger event metadata by key
	GetFromMeta(key string) string
}

// Condition automation condition interface
type Condition interface {
	// Check checks condition
	Check(trigger TriggerEvent) (bool, error)
}

// Action automation condition interface
type Action interface {
	// Execute executes action
	Execute(trigger TriggerEvent) error
}

// Automation automation interface
type Automation interface {
	// Execute automation actions
	Execute(trigger TriggerEvent) error
	// AddTrigger Add trigger to automation
	AddTrigger(trigger string)
	// GetTriggers return automation triggers list
	GetTriggers() []string
	// AddAction adds action to automation
	AddAction(action Action)
	// AddCondition set automation condition
	AddCondition(condition Condition)
}
