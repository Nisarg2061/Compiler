# Dockerfile for website service
FROM golang:1.20-alpine as builder

# Set the working directory inside the container
WORKDIR /website

# Copy the Go module files
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod tidy

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN go build -o server server.go

# Expose the port the server will run on
EXPOSE 5099

# Set the entrypoint for the container
ENTRYPOINT ["./server"]
