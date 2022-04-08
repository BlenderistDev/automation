package core

import (
	"fmt"

	"github.com/BlenderistDev/automation/interfaces"
)

type automation struct {
	actions   []interfaces.Action
	condition interfaces.Condition
	triggers  []string
}

// GetAutomation returns automation implementation
func GetAutomation() interfaces.Automation {
	return &automation{}
}

// Execute execute automation
func (a *automation) Execute(trigger interfaces.TriggerEvent) error {
	if a.condition != nil {
		checkRes, err := a.checkCondition(trigger)
		if err != nil {
			return err
		}
		if !checkRes {
			return nil
		}
	}
	err := a.executeActions(trigger)
	if err != nil {
		return err
	}
	return nil
}

// AddTrigger add trigger to automation
func (a *automation) AddTrigger(trigger string) {
	a.triggers = append(a.triggers, trigger)
}

// GetTriggers return automation triggers
func (a *automation) GetTriggers() []string {
	return a.triggers
}

// AddAction add action to automation
func (a *automation) AddAction(action interfaces.Action) {
	a.actions = append(a.actions, action)
}

// AddCondition add condition to automation
func (a *automation) AddCondition(condition interfaces.Condition) {
	a.condition = condition
}

func (a *automation) checkCondition(trigger interfaces.TriggerEvent) (bool, error) {
	res, err := a.condition.Check(trigger)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (a *automation) executeActions(trigger interfaces.TriggerEvent) error {
	for _, action := range a.actions {
		err := action.Execute(trigger)
		if err != nil {
			return fmt.Errorf("error while executing action: %s", err.Error())
		}
	}
	return nil
}
