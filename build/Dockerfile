FROM golang:1.21.3-alpine3.18 AS BuildStage
RUN apk update && apk add git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o backend ./cmd/main.go

FROM alpine:latest
WORKDIR /
COPY --from=BuildStage /app/backend backend
COPY --from=BuildStage /app/.env .env
ENTRYPOINT ["./backend"]