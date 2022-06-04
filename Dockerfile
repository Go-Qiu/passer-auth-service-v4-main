FROM golang:1.18.3-alpine3.16

WORKDIR /go

# copy the go.mod and go.sum file (from project folder) into the container /app directory
COPY go.* src/passer-auth-service-v4/
RUN cd src/passer-auth-service-v4 && go mod download

# copy all the project files into the container /app directory
COPY . src/passer-auth-service-v4/

# navigate into cmd/server folder (in the container) and build
RUN cd src/passer-auth-service-v4/cmd/server && go build -o server


# run the compiled program file
CMD ["./go/src/passer-auth-service-v4/server"]
