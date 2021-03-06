package condition

import (
	"fmt"
	"testing"

	"github.com/BlenderistDev/automation/datamapper"
	"github.com/BlenderistDev/automation/dry"
	mock_datamapper "github.com/BlenderistDev/automation/testing/datamapper"
	mock_interfaces "github.com/BlenderistDev/automation/testing/interfaces"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestEqualCondition(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mapping := make(map[string]datamapper.Mapping)

	dm := datamapper.DataMapper{Mapping: mapping}
	createdCondition := CreateEqualCondition(dm)

	switch createdCondition.(type) {
	case equalCondition:
	default:
		t.Errorf("condition type is not sendMessageAction")
	}
}

func TestEqualCondition_Check(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m1 := mock_datamapper.NewMockMapping(ctrl)
	m1.
		EXPECT().
		IsSimple().
		Return(true)

	m1.
		EXPECT().
		GetValue().
		Return("value")

	m2 := mock_datamapper.NewMockMapping(ctrl)
	m2.
		EXPECT().
		IsSimple().
		Return(true)

	m2.
		EXPECT().
		GetValue().
		Return("value")

	mapping := map[string]datamapper.Mapping{
		"value1": m1,
		"value2": m2,
	}

	trigger := mock_interfaces.NewMockTriggerEvent(ctrl)
	trigger.EXPECT().GetData().Return(make(map[string]string))
	trigger.EXPECT().GetData().Return(make(map[string]string))

	dm := datamapper.DataMapper{Mapping: mapping}
	createdCondition := CreateEqualCondition(dm)

	res, err := createdCondition.Check(trigger)
	assert.Nil(t, err)
	dry.TestCheckEqual(t, true, res)
}

func TestEqualCondition_Check_value1NotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	trigger := mock_interfaces.NewMockTriggerEvent(ctrl)
	trigger.EXPECT().GetData().Return(make(map[string]string))

	m := mock_datamapper.NewMockMapping(ctrl)

	mapping := map[string]datamapper.Mapping{
		"value2": m,
	}

	dm := datamapper.DataMapper{Mapping: mapping}
	createdCondition := CreateEqualCondition(dm)
	_, err := createdCondition.Check(trigger)
	dry.TestCheckEqual(t, fmt.Sprintf("key %s not found", "value1"), err.Error())
}

func TestEqualCondition_Check_value2NotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	trigger := mock_interfaces.NewMockTriggerEvent(ctrl)
	trigger.EXPECT().GetData().Return(make(map[string]string))
	trigger.EXPECT().GetData().Return(make(map[string]string))

	m := mock_datamapper.NewMockMapping(ctrl)
	m.
		EXPECT().
		IsSimple().
		Return(true)

	m.
		EXPECT().
		GetValue().
		Return("value")

	mapping := map[string]datamapper.Mapping{
		"value1": m,
	}

	dm := datamapper.DataMapper{Mapping: mapping}
	createdCondition := CreateEqualCondition(dm)
	_, err := createdCondition.Check(trigger)
	dry.TestCheckEqual(t, fmt.Sprintf("key %s not found", "value2"), err.Error())
}
