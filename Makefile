.PHONY: lint
lint:  ## Lint this codebase
	@go mod tidy
	@gofmt -e -s -w .
	@goimports -v -w .
	@golint .
	@go vet