package trigger

// Trigger struct for TriggerEvent implementation
type Trigger struct {
	name string
	data map[string]string
	meta map[string]string
}

// SetName sets trigger name
func (t *Trigger) SetName(name string) {
	t.name = name
}

// GetName returns trigger name
func (t *Trigger) GetName() string {
	return t.name
}

// SetData sets trigger data for specified key
func (t *Trigger) SetData(key, val string) {
	if t.data == nil {
		t.data = make(map[string]string)
	}
	t.data[key] = val
}

// GetData returns trigger data
func (t *Trigger) GetData() map[string]string {
	return t.data
}

// GetFromData returns value from trigger data for specified key
func (t *Trigger) GetFromData(key string) string {
	return t.data[key]
}

// SetMeta sets trigger metadata for specified key
func (t *Trigger) SetMeta(key, val string) {
	if t.meta == nil {
		t.meta = make(map[string]string)
	}
	t.meta[key] = val
}

// GetFromMeta returns value from trigger metadata for specified key
func (t *Trigger) GetFromMeta(key string) string {
	return t.meta[key]
}
