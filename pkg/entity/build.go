package entity

import "fmt"

// Build represents a build entity.
type Build struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

func (b Build) String() string {
	return fmt.Sprintf("{id: %s}", b.ID)
}
