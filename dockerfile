FROM golang:1.12.0-alpine3.9

# Set the working directory inside the container
WORKDIR /test

# Add the entire directory into the container
ADD . .

# Build the Go application
RUN go build -o main .
EXPOSE 8080
# Specify the command to run the executable
CMD [ "./main" ]
