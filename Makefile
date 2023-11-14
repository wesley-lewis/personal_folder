run: build
	@./bin/personal_folder

build: 
	@go build -o bin/personal_folder

test: 
	@go test -v ./...
