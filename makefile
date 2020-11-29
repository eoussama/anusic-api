all: clean start

run: 
	go run ./src/*.go

build:
	cd ./src/; go build -o ./../bin/anusic-api

start: build
	./bin/anusic-api

clean:
	rm -rf bin/ && rm -rf data/*.json
