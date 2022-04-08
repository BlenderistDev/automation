package condition

import (
	"fmt"

	"github.com/BlenderistDev/automation/interfaces"
)

type andCondition struct {
	subConditions []interfaces.Condition
}

// CreateAndCondition create and condition
func CreateAndCondition(subConditions []interfaces.Condition) (interfaces.Condition, error) {
	if len(subConditions) < 2 {
		return nil, fmt.Errorf("and condition should have at least two subconditions")
	}

	return andCondition{
		subConditions: subConditions,
	}, nil
}

// Check call check for every subcondition. If any condition is false, result is false
func (c andCondition) Check(trigger interfaces.TriggerEvent) (bool, error) {
	res := true
	for _, subCondition := range c.subConditions {
		subRes, err := subCondition.Check(trigger)
		if err != nil {
			return false, err
		}
		res = res && subRes
	}
	return res, nil
}
