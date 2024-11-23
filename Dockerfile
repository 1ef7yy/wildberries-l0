FROM golang:latest AS builder

WORKDIR /build

COPY . .
RUN go mod download && \ 
    CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/app/main.go

FROM ubuntu:22.04

WORKDIR /app


COPY --from=builder /build/main /app

CMD [ "./main" ]