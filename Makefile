all: get-deps build

build:
	@mkdir -p bin/
	@go build -o bin/benfordws benfordws.go

get-deps:
	@go get -d -v ./...

test:
	@go test -v ./...

coverage:
	@go get code.google.com/p/go.tools/cmd/cover
	@go test -cover ./...

benchmark:
	@go test -bench ./...

clean:
	@rm -rf bin/
