# Stage 1: Build the Go application
FROM golang:1.22 AS builder

WORKDIR /app

# Copy go.mod files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary (ensuring it's built in /app/)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/toungo ./src/

# Stage 2: Create a lightweight final image
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/toungo .

# Ensure the binary is executable
RUN chmod +x /root/toungo

# Expose the port (if running an HTTP server)
EXPOSE 3006

# Run the application
CMD ["./toungo"]
