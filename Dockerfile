FROM golang:1.21-alpine AS builder

# -- Set the working directory
WORKDIR /app
COPY . .

# -- Install dependencies and build the Go binary
RUN go mod tidy
RUN go build -o app main.go

# -- Use a tiny image to run the binary
FROM alpine:latest
WORKDIR /app

# -- Copy the built binary from builder
COPY --from=builder /app/app .

# -- Expose and Run
EXPOSE 8080
CMD ["./app"]
