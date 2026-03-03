BIN := html2md

.PHONY: build test clean

build:
	go build -o $(BIN) .

test:
	go test ./...

clean:
	rm -f $(BIN)
