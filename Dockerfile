FROM golang:1.24

WORKDIR /app

# Copy source code
COPY . .

# Build the application
RUN go build -o main cmd/main.go

# Expose API port
EXPOSE 8080

# Command to run the application
CMD ["./main"]