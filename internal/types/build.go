package types

import "fmt"

// Step describes a step in the running process in a build.
type Step struct {
	Name     string                 `json:"name"`
	Plugin   string                 `json:"plugin"`
	Commands []string               `json:"commands"`
	Options  map[string]interface{} `json:"options"`
	UUID     string                 `json:"uuid"`
}

// Build describes a build entity.
type Build struct {
	ID    string `json:"id,omitempty"`
	Image string `json:"image"`
	Steps []Step `json:"steps"`
}

func (t *Step) String() string {
	return fmt.Sprintf(`{
		name: %s
		plugin: %s
	}`, t.Name, t.Plugin)
}

func (b *Build) String() string {
	return fmt.Sprintf(`{
		image: %s,
		steps %v,
	}`, b.Image, b.Steps)
}
