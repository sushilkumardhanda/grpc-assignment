FROM golang:1.20.3

WORKDIR /usr/src/app
COPY . .
RUN go mod tidy

# Build the Go app
RUN go build -o main .

# Expose port 50051 to the outside world
EXPOSE 50051

# Command to run the executable
CMD ["./main"]