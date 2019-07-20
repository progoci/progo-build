FROM golang:stretch

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

CMD [ "./.docker/scripts/dev.initialize.sh" ]
