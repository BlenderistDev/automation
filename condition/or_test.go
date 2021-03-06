package condition

import (
	"fmt"
	"testing"

	"github.com/BlenderistDev/automation/dry"
	"github.com/BlenderistDev/automation/interfaces"
	mock_interfaces "github.com/BlenderistDev/automation/testing/interfaces"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestOrCondition_createOrCondition(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	subCondition1 := mock_interfaces.NewMockCondition(ctrl)
	subCondition2 := mock_interfaces.NewMockCondition(ctrl)

	subConditions := []interfaces.Condition{subCondition1, subCondition2}
	createdCondition, err := CreateOrCondition(subConditions)
	assert.Nil(t, err)

	switch condition := createdCondition.(type) {
	case orCondition:
		dry.TestCheckEqual(t, subConditions, condition.subConditions)
	default:
		t.Errorf("condition type is not orCondition")
	}
}

func TestOrCondition_createOrCondition_withLessConditions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	subCondition := mock_interfaces.NewMockCondition(ctrl)
	subConditions := []interfaces.Condition{subCondition}
	_, err := CreateOrCondition(subConditions)
	dry.TestCheckEqual(t, "or condition should have at least two subconditions", err.Error())
}

func TestOrCondition_SetConditions_checkResult(t *testing.T) {
	testOrConditionCheckWithSubCondition(t, false, false)
	testOrConditionCheckWithSubCondition(t, false, true)
	testOrConditionCheckWithSubCondition(t, true, false)
	testOrConditionCheckWithSubCondition(t, true, true)
}

func testOrConditionCheckWithSubCondition(t *testing.T, res1, res2 bool) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	trigger := mock_interfaces.NewMockTriggerEvent(ctrl)
	subCondition1 := mock_interfaces.NewMockCondition(ctrl)
	subCondition2 := mock_interfaces.NewMockCondition(ctrl)

	subCondition1.
		EXPECT().
		Check(gomock.Eq(trigger)).
		Return(res1, nil)

	subCondition2.
		EXPECT().
		Check(gomock.Eq(trigger)).
		Return(res2, nil)

	subConditions := []interfaces.Condition{subCondition1, subCondition2}
	createdCondition, err := CreateOrCondition(subConditions)
	assert.Nil(t, err)

	res, err := createdCondition.Check(trigger)
	assert.Nil(t, err)
	dry.TestCheckEqual(t, res1 || res2, res)
}

func TestOrCondition_SubConditionError(t *testing.T) {
	const errText = "some error"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	trigger := mock_interfaces.NewMockTriggerEvent(ctrl)
	subCondition1 := mock_interfaces.NewMockCondition(ctrl)
	subCondition2 := mock_interfaces.NewMockCondition(ctrl)

	subCondition1.
		EXPECT().
		Check(gomock.Eq(trigger)).
		Return(true, nil)

	subCondition2.
		EXPECT().
		Check(gomock.Eq(trigger)).
		Return(true, fmt.Errorf(errText))

	subConditions := []interfaces.Condition{subCondition1, subCondition2}
	createdCondition, err := CreateOrCondition(subConditions)
	assert.Nil(t, err)

	res, err := createdCondition.Check(trigger)
	dry.TestCheckEqual(t, false, res)
	dry.TestCheckEqual(t, errText, err.Error())
}
