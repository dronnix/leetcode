.PHONY: lint
lint:
	golangci-lint run --timeout 3m
