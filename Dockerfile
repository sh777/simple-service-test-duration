FROM golang:1.13 as builder

WORKDIR /app
COPY . /app
RUN go get ./...
RUN CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org go build -i -o app .

FROM alpine:latest

RUN apk --no-cache add ca-certificates mailcap && addgroup -S app && adduser -S app -G app
USER app
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
ENTRYPOINT ["./app"]