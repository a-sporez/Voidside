FROM golang:1.24.3

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV PORT=8080
EXPOSE 8080

RUN go build -o main .

CMD ["./main"]