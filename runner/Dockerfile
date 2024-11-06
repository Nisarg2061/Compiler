# Dockerfile for runner service
FROM golang:alpine as builder

# Set the working directory inside the container
WORKDIR /runner

# Copy the rest of the source code
COPY . .

# Download the Go dependencies
RUN ls -la

# Build the Go application
RUN go build -o runner main.go

# Expose the port the server will run on
EXPOSE 3000

# Set the entrypoint for the container
CMD ["./runner"]
