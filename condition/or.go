package condition

import (
	"fmt"

	"github.com/BlenderistDev/automation/interfaces"
)

type orCondition struct {
	subConditions []interfaces.Condition
}

// CreateOrCondition create or condition
func CreateOrCondition(subConditions []interfaces.Condition) (interfaces.Condition, error) {
	if len(subConditions) < 2 {
		return nil, fmt.Errorf("or condition should have at least two subconditions")
	}

	return orCondition{
		subConditions: subConditions,
	}, nil
}

// Check call check for every subcondition. If any condition is true, result is true
func (c orCondition) Check(trigger interfaces.TriggerEvent) (bool, error) {
	res := false
	for _, subCondition := range c.subConditions {
		subRes, err := subCondition.Check(trigger)
		if err != nil {
			return false, err
		}
		res = res || subRes
	}
	return res, nil
}
