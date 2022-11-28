.PHONY: build
build:
	@go install

.PHONY: fmt.all
fmt.all:
	@sh script/goimports.sh all