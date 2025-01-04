# Use the official Golang image as the base image
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Copy the configs folder into the container
COPY configs ./configs

# Build the Go app
RUN go build -o main ./cmd/http_server/main.go

# Expose port (change this to the port your app listens on)
EXPOSE 8080

# Run the Go app
CMD ["./main"]