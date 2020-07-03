package entity

import "fmt"

// Task describes a step in the running process in a build.
type Task struct {
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
	Tasks []Task `json:"steps"`
}

func (t *Task) String() string {
	return fmt.Sprintf(`{
		name: %s
		plugin: %s
	}`, t.Name, t.Plugin)
}

func (b *Build) String() string {
	return fmt.Sprintf(`{
		image: %s,
		steps %s,
	}`, b.Image, b.Tasks)
}
