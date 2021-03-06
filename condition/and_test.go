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

func TestAndCondition_createAndCondition(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	subCondition1 := mock_interfaces.NewMockCondition(ctrl)
	subCondition2 := mock_interfaces.NewMockCondition(ctrl)

	subConditions := []interfaces.Condition{subCondition1, subCondition2}
	createdCondition, err := CreateAndCondition(subConditions)
	assert.Nil(t, err)

	switch condition := createdCondition.(type) {
	case andCondition:
		dry.TestCheckEqual(t, subConditions, condition.subConditions)
	default:
		t.Errorf("condition type is not andCondition")
	}
}

func TestAndCondition_createAndCondition_withLessConditions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	subCondition := mock_interfaces.NewMockCondition(ctrl)
	subConditions := []interfaces.Condition{subCondition}
	_, err := CreateAndCondition(subConditions)
	dry.TestCheckEqual(t, "and condition should have at least two subconditions", err.Error())
}

func TestAndCondition_SubConditionError(t *testing.T) {
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
	createdCondition, err := CreateAndCondition(subConditions)
	assert.Nil(t, err)

	res, err := createdCondition.Check(trigger)
	dry.TestCheckEqual(t, false, res)
	dry.TestCheckEqual(t, errText, err.Error())
}

func TestAndCondition_SetConditions_checkResult(t *testing.T) {
	testAndConditionCheckWithSubCondition(t, false, false)
	testAndConditionCheckWithSubCondition(t, false, true)
	testAndConditionCheckWithSubCondition(t, true, false)
	testAndConditionCheckWithSubCondition(t, true, true)
}

func testAndConditionCheckWithSubCondition(t *testing.T, res1, res2 bool) {
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
	createdCondition, err := CreateAndCondition(subConditions)
	assert.Nil(t, err)

	res, err := createdCondition.Check(trigger)
	assert.Nil(t, err)
	dry.TestCheckEqual(t, res1 && res2, res)
}
