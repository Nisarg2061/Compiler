# Dockerfile for runner (backend API)
FROM golang:alpine

WORKDIR /app

# Install necessary packages, including Python
RUN apk add --no-cache python3 py3-pip

# Copy go mod and sum files
COPY ./go.mod ./go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 3000 for the Fiber app
EXPOSE 4000

# Run the executable
CMD ["./main"]
