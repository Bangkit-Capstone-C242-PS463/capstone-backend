FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/api

RUN go build -o /app/main .

EXPOSE 5000

ENTRYPOINT ["/app/main"]
