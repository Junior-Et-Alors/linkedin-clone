FROM golang:1.16-alpine

WORKDIR /opt/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /linkedin-clone

#FROM node:latest

#RUN chown -R node /app/node_modules
#RUN npm install -g nodemon
# File changes must be added at the very end, to avoid the installation of dependencies again and again
#RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
#COPY ../.. .

EXPOSE 8080

#CMD [ "nodemon", "--exec", "npm", "run","/linkedin-clone" ]
CMD [ "/linkedin-clone" ]