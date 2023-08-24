# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files to the working directory
COPY go.* ./

# Download and install Go dependencies
RUN go mod download

# Copy the rest of the application files to the working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080 which the application listens on
EXPOSE 8080

# Command to run the application when the container starts
CMD ["./main"]
