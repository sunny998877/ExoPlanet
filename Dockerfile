# Start from the official Golang image
FROM golang:1.18-alpine

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Expose port 8880 to the outside world
EXPOSE 8880

# Command to run the executable
CMD ["./main"]
