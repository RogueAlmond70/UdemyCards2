# Use the official Golang base image
FROM golang:1.17 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 go build -o my-card-game

# Use a minimal base image for the final image
FROM scratch

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/my-card-game ./

# Expose the port(s) that the game listens on
EXPOSE 9002

# Define the command to run the game
CMD ["./my-card-game"]
