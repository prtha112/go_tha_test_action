# Use the official Go image as the base image
FROM golang:1.20-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download


# Copy the rest of the application source code to the working directory
COPY . .

# Build the Go application
RUN go build -o main

# Create a new lightweight image for the application
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage to the current working directory
COPY --from=builder /app/main .

# Expose the port that the application listens on
EXPOSE 8095

# Run the application
CMD ["./main"]
