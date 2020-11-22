all: clean start

run: 
	go run ./src/*.go

build:
	cd ./src/; go build -o ./../bin/anusicapi

start: build
	./bin/anusicapi

clean:
	rm -rf bin/
