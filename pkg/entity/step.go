package entity

import "fmt"

// step describes a step in the running process in a build.
type step struct {
	Name     string                 `json:"name"`
	Plugin   string                 `json:"plugin"`
	Commands []string               `json:"commands"`
	Options  map[string]interface{} `json:"options"`
}

func (s step) String() string {
	return fmt.Sprintf(`{
		name: %s
		plugin: %s
	}`, s.Name, s.Plugin)
}
