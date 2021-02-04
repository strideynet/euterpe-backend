BINARY_NAME=euterpe
DATE=$(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
GIT_HASH=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell git diff --quiet || echo '-dirty')
FLAGS=-ldflags="-X 'euterpe/lib/version.BuildTime=$(DATE)' -X 'euterpe/lib/version.CommitHash=$(GIT_HASH)$(GIT_DIRTY)'"

.PHONY: build
build:
	go build -o $(BINARY_NAME) $(FLAGS) -v cmd/main.go
.PHONY: run
run:
	go run $(FLAGS) cmd/main.go service.emoji
.PHONY: test
test:
	go test -v ./...
.PHONY: clean
clean:
	go clean
	rm -rf $(BINARY_NAME)
.PHONY: protos
protos:
	./hack/protos.sh
.PHONY: mocks
mocks:
	mockery --all --keeptree