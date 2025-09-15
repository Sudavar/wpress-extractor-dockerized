# Stage 1: Builder
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . /go/src/app

RUN go mod init wpress-extractor
RUN go mod tidy
COPY inc/reader.go /go/src/app/inc
RUN go build -o wpress-extractor

# Stage 2: Slim Image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/src/app/wpress-extractor .
ENTRYPOINT ["./wpress-extractor"]
CMD []
