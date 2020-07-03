package app

import (
	"github.com/progoci/progo-core/config"
	log "github.com/sirupsen/logrus"

	"github.com/progoci/progo-build/pkg/build"
	"github.com/progoci/progo-build/pkg/database"
	"github.com/progoci/progo-build/pkg/docker"
)

// App contains dependencies used across the application.
type App struct {
	Build    *build.Build
	Config   *config.Config
	Database database.Database
	Docker   docker.Docker
	Log      *log.Logger
}
