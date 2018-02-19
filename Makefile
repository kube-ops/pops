default: check

.PHONY: check ci dependencies

dependencies:
	go get -t -v ./...

check: dependencies
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install --update
	gometalinter ./... -j2

build: dependencies
	go build ./...

ci: build check
	true
