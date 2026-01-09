FROM golang:1.22

WORKDIR /app/

RUN apt-get update && apt-get install -y librdkafka-dev

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o walletcore ./cmd/walletcore

EXPOSE 8080

CMD ["./walletcore"]