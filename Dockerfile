# Stage 1: Build the Go application
FROM golang:1.24 AS builder

WORKDIR /build

# Copy the go.mod and go.sum files to the /build directory
COPY go.mod go.sum ./
RUN go mod download

# Then copy sources and compile
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o goadv-orders cmd/dummy/dummy.go

# Stage 2: Create a lightweight final image.....
FROM alpine:latest AS runner

WORKDIR /app
COPY --from=builder /build/goadv-orders .

CMD ["/app/goadv-orders"]
