# Stage 1: Get the Air binary from the official image
FROM cosmtrek/air:latest AS air-source

# Stage 2: Build your actual environment
FROM golang:1.23-alpine

# Install git (required for go mod download)
RUN apk add --no-cache git

# Copy the air binary from the first stage to your path
COPY --from=air-source /go/bin/air /usr/local/bin/air

WORKDIR /app

# Copy dependency files first
COPY go.mod go.* ./
RUN go mod download

# Copy the rest of your code
COPY . .

# Start Air
ENTRYPOINT ["air"]