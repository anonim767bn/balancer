FROM golang:1.24.1-alpine3.21 AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOCACHE=/root/.cache/go-build

WORKDIR /app

COPY ./go.sum ./go.mod ./
COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./vendor /app/vendor

RUN --mount=type=cache,target="/root/.cache/go-build" \
    go build -mod=vendor -o app ./cmd/main.go


FROM alpine:3.21

WORKDIR /balancer

COPY --from=builder /app/app /balancer

EXPOSE 7777

CMD ["/balancer/app"]