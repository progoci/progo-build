module github.com/progoci/progo-build

go 1.14

require (
	github.com/docker/docker v1.13.1
	github.com/gin-gonic/gin v1.6.3
	github.com/gorilla/websocket v1.4.2
	github.com/pkg/errors v0.9.1
	github.com/progoci/progo-core v0.0.0-20200703210147-b5e9f8fc24ff
	github.com/progoci/progo-log v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/testify v1.4.0
	go.mongodb.org/mongo-driver v1.3.4
	google.golang.org/grpc v1.30.0
)

// For development
// replace github.com/progoci/progo-core => ./core
replace github.com/progoci/progo-log => ../log
