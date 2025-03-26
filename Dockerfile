FROM golang:1.24

WORKDIR /go/src/app

# Copy source code
COPY . .

# Expose API port
EXPOSE 8080

# Build the application
RUN go build -o main cmd/main.go

# Command to run the application
CMD ["./main"]