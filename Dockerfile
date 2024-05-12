FROM --platform=linux/amd64 golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o main ./main.go

FROM alpine
WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]