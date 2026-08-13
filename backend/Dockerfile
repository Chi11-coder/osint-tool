FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY internal ./internal
COPY cmd/ ./cmd/

# ビルド実行
RUN CGO_ENABLED=0 GOOS=linux go build -v -o ./cmd ./cmd

# --- 実行ステージ ---
FROM alpine:3.23.4

# 非ルートユーザの作成
RUN addgroup -S osint && adduser -S user -G osint
USER user

WORKDIR /app
COPY --from=builder /app/cmd .
EXPOSE 4000
CMD ["./cmd"]