package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	"github.com/progoci/progo-build/internal/app"
	"github.com/progoci/progo-build/internal/router"
	"github.com/progoci/progo-build/pkg/build"
	"github.com/progoci/progo-build/pkg/database"
	"github.com/progoci/progo-build/pkg/docker"
	"github.com/progoci/progo-build/pkg/plugin"
	"github.com/progoci/progo-build/progolog"
	"github.com/progoci/progo-core/config"
)

func main() {
	// Logger.
	logger := logrus.New()

	// Config.
	envPath, _ := filepath.Abs("./.env")
	config, err := config.New(envPath)
	if err != nil {
		logger.Fatalf("could not get configuration file")
	}

	// Database.
	database, err := getDatabase(config)
	if err != nil {
		logger.Fatalf("could not establish connection to database")
	}

	// Docker.
	dockerClient, err := docker.New(config.Get("PROXY_CONTAINER"))
	if err != nil {
		logger.Fatalf("could not establish connection to Docker daemon")
	}

	// Plugin Manager.
	pluginManager := plugin.NewManager(dockerClient)
	pluginManager.Add("Command", plugin.NewCommand())

	// Build.
	build := build.New(dockerClient, pluginManager)

	// LogClient.
	host := fmt.Sprintf("%s:%s", config.Get("PROGO_LOG_HOST"), config.Get("PROGO_LOG_PORT"))
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial log service: %v", err)
	}
	progolog.Start(conn)

	// App.
	app := &app.App{
		Config:   config,
		Build:    build,
		Database: database,
		Log:      logger,
	}

	port := fmt.Sprintf(":%s", config.Get("HOST_PORT"))

	r := gin.Default()
	r = router.BuildRoutes(r, app)

	r.Run(port)
}

// getDatabase returns a connection to the MongoDB database.
func getDatabase(config *config.Config) (*mongo.Database, error) {
	opts := &database.Opts{
		Host:     config.Get("DB_HOST"),
		Port:     config.Get("DB_PORT"),
		Database: config.Get("DB_NAME"),
		URI:      config.Get("DB_URI"),
	}

	return database.StartConnection(opts)
}
