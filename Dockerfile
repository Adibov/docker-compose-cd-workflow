FROM golang:1.23 AS base

WORKDIR /app
ADD go.mod .
RUN go mod download
ADD . .
RUN go build -o main .

FROM debian:stretch-slim

WORKDIR /app
COPY --from=base /app/main .
CMD ["/app/main"]
