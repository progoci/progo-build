package app

import (
	"github.com/progoci/progo-core/config"
	log "github.com/sirupsen/logrus"

	"github.com/progoci/progo-build/pkg/build"
	"github.com/progoci/progo-build/pkg/database"
)

// App contains dependencies used across the application.
type App struct {
	Build    *build.Manager
	Config   *config.Config
	Database *database.Database
	Log      *log.Logger
}
