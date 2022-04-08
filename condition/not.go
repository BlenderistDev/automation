package condition

import (
	"github.com/BlenderistDev/automation/interfaces"
)

type notCondition struct {
	subCondition interfaces.Condition
}

// CreateNotCondition create not condition
func CreateNotCondition(condition interfaces.Condition) interfaces.Condition {
	return notCondition{
		subCondition: condition,
	}
}

// Check return bool inversion of subcondition
func (c notCondition) Check(trigger interfaces.TriggerEvent) (bool, error) {
	res, err := c.subCondition.Check(trigger)
	if err != nil {
		return false, err
	}
	return !res, nil
}
