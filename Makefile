all: test lint vet build

build: demo

demo:
	@cd cmd/$@ && go build -o ../../bin/$@

test:
	@go test ./...

race:
	@go test -race ./...

vet:
	@go vet ./...

lint:
	@revive ./...

clean:
	@rm -rf bin
