
# Use the official Golang image to build the application
FROM golang:1.23.3 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o ./cmd/main ./cmd/.

ENV ENV prod

#  Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./cmd/main"]


