# make sure targets do not conflict with file and folder names
.PHONY: build clean test

# build the project
build:
	go build -o bin/bone

# run quality assessment checks
check:
	@echo "Running gofmt ..."
	@! gofmt -s -d -l . 2>&1 | grep -vE '^\.git/'
	@echo "Ok!"

	@echo "Running go vet ..."
	@go vet ./...
	@echo "Ok!"

	@echo "Running goimports ..."
	@! goimports -l . | grep -vF 'No Exceptions'
	@echo "Ok!"

# clean
clean:
	rm -rf bin

# format
format:
	go fmt ./...
	goimports -w .

# get all dependencies
provision:
	@echo "Getting dependencies ..."
	@go mod download
	@go get golang.org/x/tools/cmd/goimports
	@echo "Done!"

# run the binary
run:
	./bin/bone

# run tests
test:
	go test ./... -cover -v
