# Use the official Golang image as the base image
FROM golang:1.22.5

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build app
RUN go build -o main ./cmd/main.go

# Expose port 8080 from container to the host
EXPOSE 8080

# Execute main.go when the container starts
CMD ["./main"]