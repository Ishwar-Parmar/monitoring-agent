# Use the official Golang image as a base
FROM golang:1.22.3

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Navigate to monitor_metrics directory
WORKDIR /app/monitor_metrics

# Build the Go app
RUN go build -o main .
# Command to run the executable
CMD ["./main"]