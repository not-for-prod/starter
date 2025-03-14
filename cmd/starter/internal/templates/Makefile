.PHONY: linter
linter:
	golangci-lint --config .golangci.yaml run


.PHONY: generate-server
generate:
	buf dep update
	buf dep prune
	buf lint
	#buf breaking --against ".git#subdir=."
	buf generate --template buf.gen.server.yaml

.PHONY: generate-client
generate:
	buf dep update
	buf dep prune
	buf lint
	#buf breaking --against ".git#subdir=."
	buf generate --template buf.gen.client.yaml