FROM golang:latest

WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod .
COPY go.sum .

RUN go install github.com/cosmtrek/air@latest
# Download Go module dependencies
RUN go mod download

# Copy the entire source code into the Docker image
COPY . .

# Build the Go application
RUN go build -o main cmd/main.go

# Set the entry point for the container
CMD ["./main"]
