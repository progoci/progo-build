package router

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	"github.com/progoci/progo-build/internal/app"
	"github.com/progoci/progo-build/internal/types"
	"github.com/progoci/progo-build/pkg/build"
)

func createHandler(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {

		var requestBody types.Build
		c.BindJSON(&requestBody)

		buildID := primitive.NewObjectID()

		opts := &build.Opts{
			Services:          requestBody.Services,
			VirtualHostSuffix: app.Config.Get("PROXY_VIRTUAL_HOST_SUFFIX"),
			NetworkPreffix:    app.Config.Get("DOCKER_NETWORK_PREFFIX"),
			BuildID:           buildID.Hex(),
		}

		ctx := context.Background()

		buildResponse, err := app.Build.Setup(ctx, opts)
		if err != nil {
			app.Log.Errorf("Could not setup build %s: %s", buildID.Hex(), err)

			c.JSON(500, ServerErrorResponse{
				code:    500,
				message: "Could not set up build",
			})
			return
		}

		app.Log.Infof("Build %s was successfully created", buildID.Hex())
		for i, container := range buildResponse.Containers {
			app.Log.Infof("Service %s running at %s", requestBody.Services[i].Name, container.VirtualHost)
		}

		c.JSON(200, buildResponse)
	}
}
