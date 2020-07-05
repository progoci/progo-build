package build

import "github.com/progoci/progo-build/internal/types"

var availableImages = map[string]bool{
	"progoci/ubuntu18.04-php7.2-apache": true,
}

// invalidImage checks that all images in the service list are valid.
//
// It returns the first image that is not registered as valid.
func invalidImage(services []*types.Service) string {
	for _, service := range services {

		if _, ok := availableImages[service.Image]; !ok {
			return service.Image
		}

	}

	return ""
}
