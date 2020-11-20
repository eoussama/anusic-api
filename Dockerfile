# Build
FROM golang:1.14.3-alpine AS build

WORKDIR /src

COPY ./src ./

RUN go build -o /out/anusic_api .

