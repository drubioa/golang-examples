FROM golang:1.13 as builder

WORKDIR /app

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .
RUN go build hello-docker.go 

FROM alpine:latest

WORKDIR /app
COPY --from=builder app/hello-docker /app

ENTRYPOINT ["/app/hello-docker"]

EXPOSE 8080