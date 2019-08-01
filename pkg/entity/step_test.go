package entity

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStepDecoding(t *testing.T) {
	assert := assert.New(t)

	encoded := []byte(`{"name":"Step","plugin":"Shell","options":{"option1":true,"option2":false}}`)

	var s step
	e := json.Unmarshal(encoded, &s)

	if e != nil {
		assert.Fail("Could not decode json", e)
	}

	assert.Equal("Step", s.Name)
	assert.Equal("Shell", s.Plugin)
	assert.Equal(true, s.Options["option1"].(bool))
	assert.Equal(false, s.Options["option2"].(bool))
}
