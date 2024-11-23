# Use the official Golang image to create a binary.
FROM golang:1.23 as builder
WORKDIR /app

# Copy and download dependencies using go mod
COPY go.mod ./
# Commenting the go.sum line since it doesn’t exist. Might change if dependencies are added
# COPY go.sum ./
RUN go mod tidy

# Copy the source code
COPY . ./

# Build the Go app
RUN go build -o securevault

# Create a small image with just the binary
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/securevault .

# Command to run the executable
CMD ["./securevault"]

