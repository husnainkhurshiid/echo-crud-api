FROM golang:latest

WORKDIR /cmd

COPY go.mod .
COPY go.sum .
COPY cmd/*.go ./cmd/ 
RUN go mod download

RUN go build -o cmd/main

EXPOSE 8080

CMD ["./cmd/main"]
