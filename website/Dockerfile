# Dockerfile for website (frontend)
FROM golang:alpine

WORKDIR /app

# Install necessary packages, including Python
RUN apk add --no-cache python3 py3-pip

# Copy the source code
COPY . .

# Expose port 5099 for the frontend server
EXPOSE 5099

# Run the Go server
CMD ["go", "run", "server.go"]
