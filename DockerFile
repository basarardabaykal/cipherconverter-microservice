FROM golang:1.24-bookworm AS builder

WORKDIR /app

# Install git in case the build needs to fetch submodule
RUN apt-get update && apt-get install -y git && rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o microservice ./cmd/api

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/microservice .

EXPOSE 50051

CMD ["./microservice"]