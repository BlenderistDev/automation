package core

import (
	"fmt"
	"testing"

	"github.com/BlenderistDev/automation/dry"
	"github.com/BlenderistDev/automation/interfaces"
	mock_interfaces "github.com/BlenderistDev/automation/testing/interfaces"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAutomation_ExecuteNoCondition(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trigger := mock_interfaces.NewMockTriggerEvent(ctrl)

	action := mock_interfaces.NewMockAction(ctrl)
	action.
		EXPECT().
		Execute(gomock.Eq(trigger))

	actions := []interfaces.Action{action}

	automation := automation{actions: actions}

	err := automation.Execute(trigger)
	assert.Nil(t, err)
}

func TestAutomation_ExecuteManyActions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trigger := mock_interfaces.NewMockTriggerEvent(ctrl)

	action1 := mock_interfaces.NewMockAction(ctrl)
	action1.
		EXPECT().
		Execute(gomock.Eq(trigger))

	action2 := mock_interfaces.NewMockAction(ctrl)
	action2.
		EXPECT().
		Execute(gomock.Eq(trigger))

	actions := []interfaces.Action{action1, action2}

	automation := automation{actions: actions}

	err := automation.Execute(trigger)
	assert.Nil(t, err)
}

func TestAutomation_ExecuteWithConditionTrue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trigger := mock_interfaces.NewMockTriggerEvent(ctrl)

	action := mock_interfaces.NewMockAction(ctrl)
	action.
		EXPECT().
		Execute(gomock.Eq(trigger))

	condition := mock_interfaces.NewMockCondition(ctrl)
	condition.
		EXPECT().
		Check(gomock.Eq(trigger)).
		Return(true, nil)

	actions := []interfaces.Action{action}

	automation := automation{actions: actions, condition: condition}

	err := automation.Execute(trigger)
	assert.Nil(t, err)
}

func TestAutomation_ExecuteWithConditionFalse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trigger := mock_interfaces.NewMockTriggerEvent(ctrl)

	action := mock_interfaces.NewMockAction(ctrl)

	condition := mock_interfaces.NewMockCondition(ctrl)
	condition.
		EXPECT().
		Check(gomock.Eq(trigger)).
		Return(false, nil)

	actions := []interfaces.Action{action}

	automation := automation{actions: actions, condition: condition}

	err := automation.Execute(trigger)
	assert.Nil(t, err)
}

func TestAutomation_ExecuteWithConditionError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trigger := mock_interfaces.NewMockTriggerEvent(ctrl)

	action := mock_interfaces.NewMockAction(ctrl)

	condition := mock_interfaces.NewMockCondition(ctrl)
	condition.
		EXPECT().
		Check(gomock.Eq(trigger)).
		Return(true, fmt.Errorf("some error"))

	actions := []interfaces.Action{action}

	automation := automation{actions: actions, condition: condition}

	err := automation.Execute(trigger)
	dry.TestCheckEqual(t, "some error", err.Error())
}

func TestAutomation_ExecuteWithActionError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trigger := mock_interfaces.NewMockTriggerEvent(ctrl)

	actionError := "some error"
	action := mock_interfaces.NewMockAction(ctrl)
	action.
		EXPECT().
		Execute(gomock.Eq(trigger)).
		Return(fmt.Errorf(actionError))

	actions := []interfaces.Action{action}

	automation := automation{actions: actions}

	err := automation.Execute(trigger)
	dry.TestCheckEqual(t, "error while executing action: "+actionError, err.Error())
	if err != nil {
		return
	}
}

func TestAutomation_AddTrigger(t *testing.T) {
	const (
		t1 = "trigger1"
		t2 = "trigger2"
	)

	expect := []string{t1, t2}

	automation := automation{}
	automation.AddTrigger(t1)
	automation.AddTrigger(t2)
	dry.TestCheckEqual(t, expect, automation.triggers)
}

func TestAutomation_GetTriggerList(t *testing.T) {
	expect := []string{"trigger1", "trigger2"}
	automation := automation{triggers: expect}
	dry.TestCheckEqual(t, expect, automation.GetTriggers())
}

func TestAutomation_AddAction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	action := mock_interfaces.NewMockAction(ctrl)
	expected := []interfaces.Action{action}
	automation := automation{}
	automation.AddAction(action)
	dry.TestCheckEqual(t, expected, automation.actions)
}

func TestAutomation_AddCondition(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	condition := mock_interfaces.NewMockCondition(ctrl)
	automation := automation{}
	automation.AddCondition(condition)
	dry.TestCheckEqual(t, condition, automation.condition)
}

func TestGetAutomation(t *testing.T) {
	a := GetAutomation()
	switch a.(type) {
	case interfaces.Automation:
	default:
		t.Errorf("automation must implement interfaces.Automation")
	}
}
