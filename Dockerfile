FROM golang:1.17

WORKDIR /go/src/app
# VOLUME /. /go/src/app
COPY . .

# RUN go get -d -v ./...
# RUN go install -v ./...

ENTRYPOINT ["go", "run", "mycustomaction.go"]