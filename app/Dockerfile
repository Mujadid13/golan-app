# Use the official Golang image to create a build artifact.
FROM golang:1.20.5 as builder

# Create a working directory.
WORKDIR /app

# Copy go mod and sum files.
COPY app/go.mod ./go.mod
COPY app/go.sum ./go.sum

# Download all dependencies.
RUN go mod download

# Copy the source code into the container.
COPY app/. .

# Build the Go app.
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Use a Docker multi-stage build to create a lean production image.
FROM alpine:latest

# Install ca-certificates in case you need to call HTTPS endpoints.
RUN apk --no-cache add ca-certificates

# Set the working directory in the container.
WORKDIR /root/

# Copy the pre-built binary file from the previous stage.
COPY --from=builder /app/main .

# Expose port 8080 to the outside world.
EXPOSE 8080

# Command to run the executable.
CMD ["./main"]