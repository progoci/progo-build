package router

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/progoci/progo-build/internal/app"
	"github.com/progoci/progo-build/internal/types"
	builder "github.com/progoci/progo-build/pkg/build"
)

func createHandler(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {

		var requestBody types.Build
		c.BindJSON(&requestBody)

		buildID := primitive.NewObjectID()

		opts := &builder.Opts{
			Services:          requestBody.Services,
			VirtualHostSuffix: app.Config.Get("PROXY_VIRTUAL_HOST_SUFFIX"),
			NetworkPreffix:    app.Config.Get("DOCKER_NETWORK_PREFFIX"),
			BuildID:           buildID.Hex(),
		}
		ctx := context.Background()

		build, err := app.Build.Setup(ctx, opts)
		if err != nil {
			app.Log.Errorf("could not setup build %s: %s", buildID.Hex(), err)

			c.JSON(500, ServerErrorResponse{
				code:    500,
				message: "could not set up build",
			})
			return
		}

		app.Log.Infof("build %s was successfully created", buildID.Hex())
		for i, container := range build.Containers {
			app.Log.Infof("service %s running at %s", requestBody.Services[i].Name, container.VirtualHost)
		}

		id, err := app.Database.Create(build)
		if err != nil {
			app.Log.Errorf("failed to store build: %v", err)
		}

		fmt.Println(id, build.ID)

		c.JSON(200, build)
	}
}
