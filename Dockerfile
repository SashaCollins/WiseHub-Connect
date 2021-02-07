# Start from base image
FROM golang:latest

# Set the current working directory inside the container
WORKDIR /backend

RUN go get -u github.com/mattn/go-sqlite3

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy source from current directory to working directory
COPY . .

# Build the application
RUN go build -o main .

# Expose necessary port
EXPOSE 9010

# Run the created binary executable after wait for mysql container to be up
CMD ["./main"]