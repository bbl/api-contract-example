.PHONY: all clean install-deps openapi generate generate-ts-client

OPENAPI_SPEC := generated/openapi.yaml
GO_OUTPUT_DIR := generated/api
TS_CLIENT_OUTPUT_DIR := generated/ts-client

all: generate

# Install required dependencies
install-deps:
	npm install @typespec/compiler @typespec/http @typespec/openapi @typespec/openapi3
	npm install -D @hey-api/openapi-ts
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.5.1

# Generate OpenAPI spec from TypeSpec
openapi:
	@mkdir -p generated
	npx tsp compile ./contract

# Generate Go code from OpenAPI spec using oapi-codegen with strict-server
generate: openapi
	@mkdir -p $(GO_OUTPUT_DIR)
	oapi-codegen --config oapi-codegen.yaml $(OPENAPI_SPEC)

# Generate TypeScript client from OpenAPI spec
generate-ts-client: openapi
	@mkdir -p $(TS_CLIENT_OUTPUT_DIR)
	npx @hey-api/openapi-ts -f openapi-ts.config.ts

# Clean generated files
clean:
	rm -rf generated
	rm -rf node_modules
