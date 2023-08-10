FROM golang:1.18 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /go/bin/line-scrape-bot

FROM debian:buster-slim as runner

COPY --from=builder /go/bin/line-scrape-bot /go/bin/line-scrape-bot

EXPOSE 8080

CMD ["/go/bin/line-scrape-bot"]
