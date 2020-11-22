FROM golang:1.15-alpine AS build

WORKDIR .

COPY . .

RUN cd ./src/; go install

EXPOSE 8000

CMD ["$GOPATH/bin/anusic-api"]