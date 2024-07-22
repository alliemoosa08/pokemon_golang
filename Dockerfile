#Golang official Image
FROM golang:1.22-alpine AS builder

#Set current directory inside the container
WORKDIR /app

#Copy go mode and sum files
COPY go.mod go.sum ./

#Download all dependecies
RUN go mod download

#copy source code into the container
COPY . .

#Build the Golang app
RUN go build -o main .

#Start a new stage from scratch. use this step to reduce image size,
#improve security, and clean an image
FROM alpine:latest

WORKDIR /root/

#Copy the pre-built binary file from the previouse stage
COPY --from=builder /app/main .

#Expose port:8080
EXPOSE 8080

#Command to run the executable
CMD ["./main"]
