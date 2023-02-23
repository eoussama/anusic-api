# Getting a base golang image
FROM golang:1.15-alpine AS build

# Add Certificate
RUN apk --no-cache add ca-certificates

# Disable CGO
ENV CGO_ENABLED=0

# Defining the work directory
WORKDIR ./src/github.com/eoussama/anusic-api

# Copying files
COPY . .

# Building the binary
RUN go install ./src/main.go

# Renaming the binary
RUN cp /go/bin/main /go/bin/anusic-api

# Getting an empty base image
FROM scratch AS bin

# Copying the binary
COPY --from=build /go/bin/anusic-api /

# Copying the cache
COPY --from=build /go/src/github.com/eoussama/anusic-api/data /data

# Copying the deps
COPY --from=build /go/src/github.com/eoussama/anusic-api/config /config
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Exposing the port
EXPOSE 8000

# Running the API
CMD ["/anusic-api"]
