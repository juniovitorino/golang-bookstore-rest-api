FROM golang:1.21.4-alpine

RUN apk add --no-cache git
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/cosmtrek/air@latest

COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
