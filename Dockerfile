# Build stage
FROM golang:1.25-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod files if they exist
COPY go.* ./

# Download dependencies (if go.mod exists)
RUN if [ -f go.mod ]; then go mod download; fi

# Copy source code and data file
COPY main.go .
COPY reasons.json .

# Build the application
# CGO_ENABLED=0 for static binary
# -ldflags="-w -s" strips debug info for smaller binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o app main.go

# Final stage - minimal image
FROM scratch

# Copy the binary from builder
COPY --from=builder /app/app /app

# Copy the data file
COPY --from=builder /app/reasons.json /reasons.json

# Run the application
ENTRYPOINT ["/app"]