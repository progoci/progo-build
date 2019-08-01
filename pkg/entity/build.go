package entity

import "fmt"

// Build describes a build entity.
type Build struct {
	ID    string `json:"id,omitempty"`
	Image string `json:"image"`
	Steps []step `json:"steps"`
}

func (b Build) String() string {
	return fmt.Sprintf(`{
		image: %s,
		steps %s,
	}`, b.Image, b.Steps)
}
