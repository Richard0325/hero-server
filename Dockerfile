FROM golang:1.17-alpine

WORKDIR /hero-server

COPY . /hero-server/

RUN go mod tidy

ENTRYPOINT ["go", "run", "cmd/server/main.go"]