M golang:1.21-alpine

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o receipt-processor

# Expose port (assuming you'll use 8080)
EXPOSE 8080

# Run the executable
CMD ["./receipt-processor"]
