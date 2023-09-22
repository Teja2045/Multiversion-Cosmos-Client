.PHONY: build

include .env
export

build:
	@echo "deleting previous go.mod file..."
	@rm -rf go.mod
	@rm -rf go.sum
	@echo "successfully deleted!"
	@echo "Building..."
	@cp ./go-mods/gomod$(SDK_VERSION) ./go.mod
	@cp ./go-mods/gosum$(SDK_VERSION) ./go.sum
	@echo "CosmosSDK v$(SDK_VERSION) version build is successful!"

run:
	@go run -tags "$(SDK_VERSION) server" .