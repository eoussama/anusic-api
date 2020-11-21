run: 
	go run ./src/*.go

build:
	cd ./src/; go build -o ./../bin/anusicapi

clean:
	rm -rf bin/
