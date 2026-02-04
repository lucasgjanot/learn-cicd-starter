# Stage 1 — build
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$(go env GOARCH) go build -o notely


FROM debian:stable-slim
RUN apt-get update && apt-get install -y ca-certificates
COPY --from=builder /app/notely /usr/bin/notely

CMD ["notely"]
