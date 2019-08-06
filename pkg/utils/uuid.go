package utils

import (
	"github.com/google/uuid"
)

// GetUUID returns a new randomly-generated UUID value.
func GetUUID() string {
	uuid := uuid.New()

	return uuid.String()
}
