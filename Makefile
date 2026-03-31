# Makefile
.PHONY: generate

generate:
	oapi-codegen --config oapi-codegen.yaml api/openapi.yaml
	@echo "✓ Go types and server interfaces generated"