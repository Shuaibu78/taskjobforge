# Stage 1: build
FROM golang:1.24.1-alpine AS builder
WORKDIR /app

# Copy go.mod first (go.sum might not exist yet)
COPY go.mod ./
# Generate go.sum and download dependencies
RUN go mod download && go mod verify

# Copy source
COPY . .

# Build binary
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o /taskjobforge ./cmd/worker

# Stage 2: minimal runtime
FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=builder /taskjobforge /taskjobforge

EXPOSE 8080
USER 65532:65532
CMD ["/taskjobforge"]