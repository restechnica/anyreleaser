# make sure targets do not conflict with file and folder names
.PHONY: build clean test

# build the project
build:
	go build -o bin/bone

# clean
clean:
	rm -rf bin

# run the binary
run:
	./bin/bone

# run tests
test:
	go test ./... -cover -v
