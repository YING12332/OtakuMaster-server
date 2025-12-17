# ===== Build stage =====
FROM golang:1.25-alpine AS builder

WORKDIR /app

# 先拷依赖文件，提高缓存命中
COPY go.mod go.sum ./
RUN go mod download

# 再拷源码
COPY . .

# 构建静态二进制
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/server

# ===== Runtime stage =====
FROM alpine:3.20

WORKDIR /app
COPY --from=builder /app/server /app/server

EXPOSE 8080
CMD ["/app/server"]
