# Use the official Golang image to create a binary.
FROM golang:1.23 as builder

# Create a non-root user and switch to it
RUN useradd -ms /bin/bash secureuser
WORKDIR /home/secureuser/app

# Copy files as root, then change ownership to the non-root user
COPY go.mod ./
RUN chown -R secureuser:secureuser /home/secureuser/app

# Switch to the non-root user
USER secureuser

# Download dependencies
RUN go mod tidy

# Copy the source code
COPY . ./

# Build the Go app, disabling VCS stamping
RUN go build -buildvcs=false -o securevault

# Create a small image with just the binary
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /home/secureuser/app/securevault .

# Command to run the executable
CMD ["./securevault"]

