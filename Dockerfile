# Build
FROM golang:1.14.3-alpine AS build

WORKDIR /src

COPY . .

RUN go build -o /out/example .

# Run
FROM scratch AS bin

COPY --from=build /out/anusic

CMD ["npm", "start"]
