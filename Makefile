default: out/example

clean:
	rm -rf out

test: *.go
	go test implementation.go implementation_test.go
	go test implementation.go handler.go handler_test.go

build:
	mkdir -p out
	go build -o out/example ./cmd/example/main.go
