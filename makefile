BINARY_NAME=gmm
DATE=$(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
GIT_HASH=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell git diff --quiet || echo '-dirty')
FLAGS=-ldflags="-X 'gmm/lib/version.BuildTime=$(DATE)' -X 'gmm/lib/version.CommitHash=$(GIT_HASH)$(GIT_DIRTY)'"

build:
	go build -o $(BINARY_NAME) $(FLAGS) -v cmd/main.go
run:
	go run $(FLAGS) cmd/main.go service.emoji
test:
	go test -v ./...
clean:
	go clean
	rm -rf $(BINARY_NAME)