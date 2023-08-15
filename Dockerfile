
# Use an official Golang runtime as the base image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download and install dependencies
RUN go mod download  && go mod verify

# Copy the entire project into the container
COPY . .

# Build the Go application
RUN go build -o go .

# Expose the port the application runs on
EXPOSE 3001

# Command to run the application
CMD ["./go"]
