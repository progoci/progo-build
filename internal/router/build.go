package router

import (
	"github.com/gin-gonic/gin"
	"github.com/progoci/progo-build/internal/app"
)

// BuildRoutes returns the routes for build
func BuildRoutes(r *gin.Engine, app *app.App) *gin.Engine {
	r.POST("/", createHandler(app))

	return r
}
