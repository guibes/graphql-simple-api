FROM golang:1.16.5-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/graphql-simple-api

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .


# Build the Go app
RUN CGO_ENABLED=0 go build -o ./out/graphql-simple-api .

# Start fresh from a smaller image
FROM alpine:3.9 
RUN apk add ca-certificates

COPY --from=build_base /tmp/graphql-simple-api/out/graphql-simple-api /app/graphql-simple-api

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["/app/out/graphql-simple-api"]