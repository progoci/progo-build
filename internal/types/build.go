package types

// Step describes a step in the running process in a build.
type Step struct {
	Name     string                 `json:"name"`
	Plugin   string                 `json:"plugin"`
	Commands []string               `json:"commands"`
	Options  map[string]interface{} `json:"options"`
}

// Service describes a service for a build.
type Service struct {
	Name  string  `json:"name"`
	Image string  `json:"image"`
	Steps []*Step `json:"steps"`
}

// Build describes a build entity.
type Build struct {
	ID       string     `json:"id,omitempty"`
	Services []*Service `json:"services"`
}
