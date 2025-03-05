FROM golang:alpine as builder
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint...@latest
RUN apk add --no-cache make musl-dev gcc g++ bash
WORKDIR /app
RUN go build -o service main.go

FROM alpine:latest as production
COPY --from=builder /service /service