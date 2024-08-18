FROM golang:1.19 as builder

WORKDIR /app

# Go modulesをダウンロード
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# Playwrightのインストールを実行して、ブラウザと依存関係をインストール
RUN go install github.com/playwright-community/playwright-go/cmd/playwright@latest && \
    playwright install --with-deps

# アプリケーションをビルド
RUN go build -o /go/bin/line-scrape-bot

# 実行環境の設定
FROM debian:buster-slim as runner

# OSの必要なパッケージをインストール
RUN apt-get update \
    && apt-get install -y ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# ビルド成果物をコピー
COPY --from=builder /go/bin/line-scrape-bot /go/bin/line-scrape-bot
COPY --from=builder /app/config /app/config

# 必要なPlaywrightのブラウザをコピー
COPY --from=builder /root/.cache/ms-playwright /root/.cache/ms-playwright

EXPOSE 8080

CMD ["/go/bin/line-scrape-bot"]
