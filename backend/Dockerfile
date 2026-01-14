FROM golang:1.23-alpine
# install Air for hot reloading
RUN go install github.com/air-verse/air@last 
# set working dir
WORKDIR /app
# copy go mod files for faster builds
COPY  go.mod go.sum ./
RUN go mod download
# Copy the rest of the application
COPY . . 
# Use Air as the entry point for development
CMD ["air", "-c", ".air.toml"] 
