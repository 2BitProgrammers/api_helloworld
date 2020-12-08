## To build:  docker build . -t 2bitprogrammers/api_helloworld

## Use golang image to build executable
FROM golang:alpine AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
COPY $PWD/src/go.mod .
COPY $PWD/src/main.go .
RUN go mod download
RUN go build -o api_helloworld . 


## Build final image from scratch (copy executeable into empty container)
FROM scratch 
WORKDIR /
COPY --from=builder /build/api_helloworld . 
ENTRYPOINT [ "/api_helloworld" ]