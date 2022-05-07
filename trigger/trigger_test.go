package trigger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrigger_SetName(t *testing.T) {
	const name = "name"
	trigger := Trigger{}
	trigger.SetName(name)
	assert.Equal(t, name, trigger.name)
}

func TestTrigger_GetName(t *testing.T) {
	const name = "name"
	trigger := Trigger{name: name}
	assert.Equal(t, name, trigger.GetName())
}

func TestTrigger_SetData(t *testing.T) {
	const key = "key"
	const value = "value"
	data := map[string]string{
		key: value,
	}
	trigger := Trigger{}
	trigger.SetData(key, value)
	assert.Equal(t, data, trigger.data)
}

func TestTrigger_GetData(t *testing.T) {
	const key = "key"
	const value = "value"
	data := map[string]string{
		key: value,
	}
	trigger := Trigger{data: data}
	assert.Equal(t, data, trigger.GetData())
}

func TestTrigger_GetFromData(t *testing.T) {
	const key = "key"
	const value = "value"
	data := map[string]string{
		key: value,
	}
	trigger := Trigger{data: data}
	assert.Equal(t, value, trigger.GetFromData(key))
}

func TestTrigger_SetMeta(t *testing.T) {
	const key = "key"
	const value = "value"
	meta := map[string]string{
		key: value,
	}
	trigger := Trigger{}
	trigger.SetMeta(key, value)
	assert.Equal(t, meta, trigger.meta)
}

func TestTrigger_GetFromMeta(t *testing.T) {
	const key = "key"
	const value = "value"
	meta := map[string]string{
		key: value,
	}
	trigger := Trigger{meta: meta}
	assert.Equal(t, value, trigger.GetFromMeta(key))
}
