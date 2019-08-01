package entity

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildDecoding(t *testing.T) {
	assert := assert.New(t)

	encoded := []byte(`{
		"image": "progoci/ubuntu-18.04:php-7.2",
		"steps": [
			{
				"name": "Step 1",
				"plugin": "Shell",
				"commands": [
					"ls -al",
					"cd path"
				]
			},
			{
				"name": "Step 2",
				"plugin": "Drupal",
				"options": {
					"option1": true,
					"option2": false
				}
			}
		]
	}`)

	var b Build
	e := json.Unmarshal(encoded, &b)

	if e != nil {
		assert.Fail("Could not decode json", e)
	}

	assert.Equal("progoci/ubuntu-18.04:php-7.2", b.Image)
	assert.Equal("Step 1", b.Steps[0].Name)
	assert.Equal("Shell", b.Steps[0].Plugin)
	assert.Equal("ls -al", b.Steps[0].Commands[0])
	assert.Equal(true, b.Steps[1].Options["option1"].(bool))
}
