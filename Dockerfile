# Golang Image:Tag
FROM golang:1.21

# Set the working directory
WORKDIR /go/src/app

# Copy local ke container docker
COPY go.mod .
COPY main.go .

# Build golang
RUN go get
RUN go build -o bin .

# Expose the port the app runs on
EXPOSE 8081

# Define the command to run your app
CMD ["app"]
