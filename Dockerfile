FROM golang:1.14-stretch

RUN mkdir /app

COPY . /app

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -mod=readonly -o server" -command="./server"
