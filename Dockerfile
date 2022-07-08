FROM golang:1.16-alpine

WORKDIR /opt/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /linkedin-clone

EXPOSE 8080

CMD [ "/linkedin-clone" ]