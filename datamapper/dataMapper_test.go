package datamapper

import (
	"math"
	"strconv"
	"testing"

	"github.com/BlenderistDev/automation/dry"
	mock_datamapper "github.com/BlenderistDev/automation/testing/datamapper"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetFromMap_simpleMapping(t *testing.T) {
	const name = "name"
	const value = "value"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mapping := getMapping(ctrl, true, name, value)
	datamapper := DataMapper{Mapping: mapping}

	mapValue, err := datamapper.GetFromMap(make(map[string]string), "name")
	assert.Nil(t, err)
	dry.TestCheckEqual(t, value, mapValue)
}

func TestGetFromMap_notSimpleMapping(t *testing.T) {
	const value = "value"
	const name = "name"
	const resultValue = "test_value"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mapping := getMapping(ctrl, false, name, value)
	datamapper := DataMapper{Mapping: mapping}

	data := map[string]string{
		"value": resultValue,
	}

	mapValue, err := datamapper.GetFromMap(data, name)
	assert.Nil(t, err)
	dry.TestCheckEqual(t, resultValue, mapValue)
}

func TestGetFromInt32_simpleMapping(t *testing.T) {
	const name = "name"
	const value = "123"
	var valueInt int32 = 123

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mapping := getMapping(ctrl, true, name, value)
	datamapper := DataMapper{Mapping: mapping}

	data := make(map[string]string)
	mapValue, err := datamapper.GetFromMapInt32(data, "name")
	assert.Nil(t, err)
	dry.TestCheckEqual(t, valueInt, mapValue)
}

func TestGetFromInt32_bigValue(t *testing.T) {
	const name = "name"
	const valueInt = math.MaxInt32 + 1

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mapping := getMapping(ctrl, true, name, strconv.Itoa(valueInt))
	datamapper := DataMapper{Mapping: mapping}

	data := make(map[string]string)
	_, err := datamapper.GetFromMapInt32(data, name)
	dry.TestCheckEqual(t, "number 2147483648 is greater, than MaxInt32", err.Error())
}

func TestGetFromMapInt32_notSimpleMapping(t *testing.T) {
	const value = "value"
	const name = "name"
	const resultValue = "123"
	var resultValueInt int32 = 123

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mapping := getMapping(ctrl, false, name, value)
	datamapper := DataMapper{Mapping: mapping}

	data := map[string]string{
		"value": resultValue,
	}
	mapValue, err := datamapper.GetFromMapInt32(data, name)
	assert.Nil(t, err)
	dry.TestCheckEqual(t, resultValueInt, mapValue)
}

func TestGetFromInt64_simpleMapping(t *testing.T) {
	const name = "name"
	const value = "123"
	var valueInt int64 = 123

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mapping := getMapping(ctrl, true, name, value)
	datamapper := DataMapper{Mapping: mapping}

	data := make(map[string]string)
	mapValue, err := datamapper.GetFromMapInt64(data, name)
	assert.Nil(t, err)
	dry.TestCheckEqual(t, valueInt, mapValue)
}

func TestGetFromMapInt64_notSimpleMapping(t *testing.T) {
	const value = "value"
	const name = "name"
	const resultValue = "123"
	const resultValueInt = 123

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mapping := getMapping(ctrl, false, name, value)
	datamapper := DataMapper{Mapping: mapping}

	data := map[string]string{
		"value": resultValue,
	}
	mapValue, err := datamapper.GetFromMapInt64(data, name)
	assert.Nil(t, err)

	if mapValue != resultValueInt {
		t.Errorf("expected: %d, actual: %d", resultValueInt, mapValue)
	}
}

func TestGetFromInt64_valueNotExist(t *testing.T) {
	const key = "name"

	mapping := make(map[string]Mapping)
	datamapper := DataMapper{Mapping: mapping}

	data := make(map[string]string)
	_, err := datamapper.GetFromMapInt64(data, key)
	dry.TestCheckEqual(t, "key "+key+" not found", err.Error())
}

func TestGetFromInt64_valueIncorrect(t *testing.T) {
	const name = "key"
	const value = "test"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mapping := getMapping(ctrl, true, name, value)
	datamapper := DataMapper{Mapping: mapping}

	data := make(map[string]string)
	_, err := datamapper.GetFromMapInt64(data, name)
	dry.TestCheckEqual(t, "strconv.ParseInt: parsing \""+value+"\": invalid syntax", err.Error())
}

func TestGetFromMap_valueNotExist(t *testing.T) {
	const key = "name"

	mapping := make(map[string]Mapping)
	datamapper := DataMapper{Mapping: mapping}

	data := make(map[string]string)
	_, err := datamapper.GetFromMap(data, key)
	dry.TestCheckEqual(t, "key "+key+" not found", err.Error())
}

func TestGetFromMap_notSimpleMapping_valueNotExist(t *testing.T) {
	const name = "name"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mapping := getMapping(ctrl, false, name, "")
	datamapper := DataMapper{Mapping: mapping}

	data := make(map[string]string)
	_, err := datamapper.GetFromMap(data, name)
	dry.TestCheckEqual(t, "key "+name+" not found in trigger data", err.Error())
}

func TestGetFromBool_simpleMapping_trueValue(t *testing.T) {
	testGetFromBool_simpleMapping(t, "123", true)
}

func TestGetFromBool_simpleMapping_falseValue(t *testing.T) {
	testGetFromBool_simpleMapping(t, "", false)
}

func TestGetFromMapBool_valueNotExist(t *testing.T) {
	const key = "name"

	mapping := make(map[string]Mapping)
	datamapper := DataMapper{Mapping: mapping}

	data := make(map[string]string)
	_, err := datamapper.GetFromMapBool(data, key)

	dry.TestCheckEqual(t, "key "+key+" not found", err.Error())
}

func testGetFromBool_simpleMapping(t *testing.T, value string, res bool) {
	const name = "name"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mapping := getMapping(ctrl, true, name, value)
	datamapper := DataMapper{Mapping: mapping}

	data := make(map[string]string)
	mapValue, err := datamapper.GetFromMapBool(data, name)
	assert.Nil(t, err)
	dry.TestCheckEqual(t, res, mapValue)
}

func getMapping(ctrl *gomock.Controller, isSimple bool, name, value string) map[string]Mapping {
	m := mock_datamapper.NewMockMapping(ctrl)
	m.
		EXPECT().
		IsSimple().
		Return(isSimple)

	m.
		EXPECT().
		GetValue().
		Return(value)

	mapping := map[string]Mapping{
		name: m,
	}
	return mapping
}
