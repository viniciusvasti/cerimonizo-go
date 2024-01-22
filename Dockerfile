FROM golang:1.21.6-alpine3.19 AS base
RUN apk update
WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN go build -o cerimonizo ./cmd/main.go

FROM alpine:3.16.5 AS binary
COPY --from=base /app/cerimonizo .
EXPOSE 3000
CMD ["./cerimonizo"]