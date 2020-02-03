FROM golang:1.13 as builder

WORKDIR /go/src/github.com/sh777/simple-service-test-duration
COPY . /go/src/github.com/sh777/simple-service-test-duration
RUN CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org go build -i -o app .

FROM alpine:latest

RUN apk --no-cache add ca-certificates mailcap && addgroup -S app && adduser -S app -G app
USER app
WORKDIR /app
COPY --from=builder /go/src/github.com/sh777/simple-service-test-duration/app .
EXPOSE 8080
ENTRYPOINT ["./app"]
