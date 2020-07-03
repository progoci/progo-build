module github.com/progoci/progo-build

go 1.14

require (
	9fans.net/go v0.0.2
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.13.1
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/gorilla/websocket v1.4.2
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/progoci/progo-core v0.0.0-20200620225418-29f5f579b134
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/testify v1.6.1
	go.mongodb.org/mongo-driver v1.3.4
)

// For development
replace github.com/progoci/progo-core => ../../core
