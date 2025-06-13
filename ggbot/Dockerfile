FROM golang:1.21-alpine

WORKDIR /bot

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["go", "run", "main.go"]