all: clean start

run: 
	go run ./src/*.go

build:
	cd ./src/; go build -o ./../bin/anusic-api

start: build
	./bin/anusic-api

start: build
	./bin/anusicapi

clean:
	rm -rf bin/
