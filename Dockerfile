FROM golang:1.25.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . /opt/app/

WORKDIR /opt/app/

RUN GOOS=linux go build -o /serv cmd/main.go 

ENTRYPOINT ["/serv"]