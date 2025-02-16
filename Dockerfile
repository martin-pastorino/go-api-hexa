
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
RUN CGO_ENABLED=0 GOOS=linux go build  -a -installsuffix cgo -o ./cmd/main ./cmd/.


FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /root/

COPY --from=builder /app .

ENV ENV prod

EXPOSE 8080

CMD ["./cmd/main"]
