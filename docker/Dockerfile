# Getting a base golang image
FROM golang:1.15-alpine AS build

# Defining the work directory
WORKDIR .

# Copying files
COPY . .

# Building the binary
RUN cd ./src/; go install

# Exposing the port
EXPOSE 8000

# Running the API
CMD ["./bin/anusic-api"]
