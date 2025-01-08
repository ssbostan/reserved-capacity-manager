FROM golang:1.23-alpine AS builder

WORKDIR /opt/src

COPY . .

RUN go build cmd/consumer.go

FROM alpine:3

WORKDIR /opt/app

COPY --from=builder /opt/src/consumer .

RUN adduser -DH rootless

USER 1000

ENTRYPOINT ["/opt/app/consumer"]

CMD ["-workers", "2", "-memory", "512"]
