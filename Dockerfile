FROM golang:1.18 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /go/bin/line-scrape-bot

FROM debian:buster-slim as runner

RUN apt-get update \
    && apt-get install -y ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /go/bin/line-scrape-bot /go/bin/line-scrape-bot

EXPOSE 8080

CMD ["/go/bin/line-scrape-bot"]
