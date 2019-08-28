package entity

import "fmt"

// Task describes a step in the running process in a build.
type Task struct {
	Name     string                 `json:"name"`
	Plugin   string                 `json:"plugin"`
	Commands []string               `json:"commands"`
	Options  map[string]interface{} `json:"options"`
	UUID     string									`json:"uuid"`
}

func (t Task) String() string {
	return fmt.Sprintf(`{
		name: %s
		plugin: %s
	}`, t.Name, t.Plugin)
}
