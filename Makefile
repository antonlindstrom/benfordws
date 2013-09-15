all:
	@mkdir -p bin/
	@go build -o bin/benfordws benfordws.go

get-deps:
	@go get -d -v ./...

test:
	@go test -v ./...

clean:
	@rm -rf bin/
