FROM golang:1.24.3

WORKDIR /bot

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["go", "run", "main.go"]