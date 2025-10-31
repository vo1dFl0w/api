FROM golang:1.24.3-alpine AS builder

WORKDIR /api

RUN apk --no-cache add bash git make gcc gettext musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

ENV CGO_ENABLED=0
RUN go build -ldflags="-s -w" -o api ./cmd/api/main.go

FROM alpine AS runner
RUN apk add --no-cache ca-certificates

WORKDIR /api
COPY --from=builder /api/api /api/api
COPY --from=builder /api/config.env /api/config.env

EXPOSE 8080

CMD ["./api"]