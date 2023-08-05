FROM golang:alpine AS builder
WORKDIR /swampy-service
COPY  . .
RUN apk update
RUN CGO_ENABLED=0 GOOS=linux go build -o main -ldflags="-s -w" -a -installsuffix cgo ./cmd/main.go

FROM alpine
WORKDIR /swampy-service
COPY --from=builder ./swampy-service .

EXPOSE 8080

ENTRYPOINT ["./main"]
