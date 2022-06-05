# Build stage
FROM golang:1.18.3-alpine3.16 AS builder
WORKDIR /app

# copy the go.mod and go.sum file (from project folder) into the container /app directory
COPY go.* ./
RUN go mod download

# copy all the project files into the container /app directory
COPY . ./

# navigate into cmd/server folder (in the container) and build
RUN cd cmd/server && go build -o server

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/cmd/server/server .

EXPOSE 9091

# run the compiled program file
CMD ["/app/server"]
