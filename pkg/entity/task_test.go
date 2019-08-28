package entity

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStepDecoding(t *testing.T) {
	assert := assert.New(t)

	encoded := []byte(`{"name":"Step","plugin":"Shell","options":{"option1":true,"option2":false}}`)

	var tsk Task
	e := json.Unmarshal(encoded, &tsk)

	if e != nil {
		assert.Fail("Could not decode json", e)
	}

	assert.Equal("Step", tsk.Name)
	assert.Equal("Shell", tsk.Plugin)
	assert.Equal(true, tsk.Options["option1"].(bool))
	assert.Equal(false, tsk.Options["option2"].(bool))
}
