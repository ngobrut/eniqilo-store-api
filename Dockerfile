FROM golang:1.22

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
COPY .env .

RUN GOOS=linux GOARCH=amd64 go build -o /main

EXPOSE 8080
CMD ["/main"]