# ===== builder =====
FROM golang:1.25-alpine AS builder
WORKDIR /app

# 可选：如果你项目里用了 CGO（一般没有），需要 gcc/musl-dev
# RUN apk add --no-cache build-base

COPY . .
RUN go build -mod=vendor -o server ./cmd/server

# ===== runtime =====
FROM alpine:3.20
WORKDIR /app

# 证书（如果后面要 https/oss/minio 会用到）
RUN apk add --no-cache ca-certificates && update-ca-certificates

COPY --from=builder /app/server /app/server

EXPOSE 8080
ENTRYPOINT ["/app/server"]
