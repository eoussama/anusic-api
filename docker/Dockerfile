# Getting a base golang image
FROM golang:1.15-alpine AS build

# Disable CGO
ENV CGO_ENABLED=0

# Defining the work directory
WORKDIR ./src

# Copying files
COPY ./src .
COPY ./go.mod .
COPY ./data ./data
COPY ./config ./config

# Building the binary
RUN go install

# Getting an empty base image
FROM scratch AS bin

# Copying the binary
COPY --from=build /go/bin/anusic-api /
COPY --from=build /go/src/data /data
COPY --from=build /go/src/config /config

# Exposing the port
EXPOSE 8000

# Running the API
# CMD ["/go/bin/anusic-api"]
CMD ["/anusic-api"]
