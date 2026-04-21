APP_NAME := my-go-project
DIST_DIR := dist
BIN := $(DIST_DIR)/$(APP_NAME)

.PHONY: build test test-smoke clean

build:
	mkdir -p $(DIST_DIR)
	go build -o $(BIN) .

test:
	go test -v ./...

# Smoke test: run a small, critical subset to quickly verify core behavior.
test-smoke:
	go test -v -run 'TestCastSuite|TestToStringE' ./...

clean:
	rm -rf $(DIST_DIR)
