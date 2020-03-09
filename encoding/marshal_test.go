package encoding

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshal1(t *testing.T) {
	data := []byte(`{"type":1}`)
	payload, err := json.Marshal(data)
	assert.Nil(t, err)
	assert.NotNil(t, payload)
}

func TestMarshalMarshalWithInvalidJSON(t *testing.T) {
	data := []byte(`{"type":1`)
	payload, err := json.Marshal(data)

	assert.Error(t, err)
	assert.Nil(t, payload)
}

func TestUnmarshalMarshalWithInvalidJSON(t *testing.T) {
	var payload json.RawMessage
	data := []byte(`{"type":1`)

	err := json.Unmarshal(data, &payload)

	assert.Error(t, err)
	assert.Nil(t, payload)
}
