.PHONY: test
test:
	go test ./... -cover

.PHONY: run-build
run-build:
	go run cmd/ds/main.go
