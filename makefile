.PHONY: all build install clean

# build into ./bin/tree-top
build:
	go build -o /bin/tree-top ./cmd/tree-top

install:
	go install github.com/faizalam/tree-top/cmd/tree-top@latest

clean:
	rm -rf bin