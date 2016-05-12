all:
	cd pili; go build -v

test:
	cd pili; go test -v

govet:
	@find . -name '*.go' | xargs -L 1 go tool vet

gofmt:
	@test `gofmt -l . | wc -l` -eq 0
