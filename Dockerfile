
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
RUN go build -o main ./cmd/.

#  Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

# # Etapa 1: Build
# FROM golang:1.23.3 AS builder

# # Establecer el directorio de trabajo
# WORKDIR /app

# # Copiar los archivos necesarios para las dependencias
# COPY go.mod go.sum ./

# # Descargar las dependencias
# RUN go mod download

# # Copiar todo el proyecto (necesario para compilar)
# COPY . .

# # Construir el binario
# RUN go build -o main ./cmd/.

# # Etapa 2: Imagen final
# FROM alpine:latest

# # Establecer el directorio de trabajo en la nueva imagen
# WORKDIR /root/

# # Copiar el binario desde la etapa builder
# COPY --from=builder /app/main .

# # Exponer el puerto necesario
# EXPOSE 8080

# # Ejecutar el binario al iniciar el contenedor
# CMD ["./main"]
