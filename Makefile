.PHONY: build
build:
	@go install

.PHONY: fmt.all
fmt.all:
	@sh scripts/goimports.sh all