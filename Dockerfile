FROM golang:stretch

RUN go get github.com/pilu/fresh

CMD [ "./.docker/scripts/dev.initialize.sh" ]
