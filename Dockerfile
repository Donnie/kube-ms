FROM golang:1.14-alpine AS builder
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GIN_MODE=release

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Run test
RUN go test ./...

# Build the application
RUN go build -o main .

############################
# STEP 2 build a small image
############################
FROM busybox

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/main /

EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["/main"]
