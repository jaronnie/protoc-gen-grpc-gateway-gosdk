.PHONY: build
build:
	@go install ./cmd/protoc-gen-go-httpsdk

.PHONY: fmt.all
fmt.all:
	@sh scripts/goimports.sh all