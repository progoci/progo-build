package entity

// Container describes a new container.
type Container struct {
	ID   string // The container ID (assigned at creation by Docker).
	Host string // The host where the live site is accessible from.
}
