FROM golang:1.16-alpine as build

WORKDIR /opt/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /opt/app .

FROM golang:1.16-alpine as run

WORKDIR /opt/app
RUN go get github.com/codegangsta/gin
ENV BIN_APP_PORT=8080
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

ENTRYPOINT [ "gin","run","main.go" ]