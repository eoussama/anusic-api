run: 
	go run ./src/*.go

build:
	cd ./src/; go build -o anusicapi

clean:
	rm -rf src/anusicapi
