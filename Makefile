
HELLO_FILES := $(shell find ./serverless/hello -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: all clean serverless frontend

all: serverless frontend

serverless: 
	mkdir -p functions
	go get -v -d ./...
	go build -o ./functions/hello $(HELLO_FILES)

frontend:
	npm run --prefix frontend build

clean:
	rm -rf ./functions
	rm -rf ./frontend/dist
