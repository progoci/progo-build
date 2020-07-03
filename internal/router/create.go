package router

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	"github.com/progoci/progo-build/internal/app"
	"github.com/progoci/progo-build/internal/entity"
	"github.com/progoci/progo-build/pkg/build"
)

func createHandler(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {

		var requestBody entity.Build
		c.BindJSON(&requestBody)

		buildID := primitive.NewObjectID()

		opts := &build.Opts{
			Image:             requestBody.Image,
			VirtualHostSuffix: app.Config.Get("PROXY_VIRTUAL_HOST_SUFFIX"),
			NetworkPreffix:    app.Config.Get("DOCKER_NETWORK_PREFFIX"),
			BuildID:           buildID.Hex(),
		}

		buildResponse, err := app.Build.Setup(context.Background(), opts)
		if err != nil {
			app.Log.Errorf("Could not create or start container: %s", err)

			c.JSON(500, ServerErrorResponse{
				code:    500,
				message: "Could not create container",
			})
			return
		}

		app.Log.Infof("Build running at %s", buildResponse.VirtualHost)

		c.JSON(200, buildResponse)
	}
}
