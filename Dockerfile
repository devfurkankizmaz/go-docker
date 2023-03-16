# Golang latest image
FROM golang:latest

# Set app as current working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies required
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside docker container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080. Application is running at 8080, so we need to expose that port
EXPOSE 8080

# Run the application
CMD ["./main"]