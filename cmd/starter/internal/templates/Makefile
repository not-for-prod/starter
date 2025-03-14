.PHONY: linter
linter:
	golangci-lint --config .golangci.yaml run

# runs buf generate with installation of all deps
.PHONY: proto
proto:
	buf dep update
	buf dep prune
	buf lint
	#buf breaking --against ".git#subdir=."
	buf generate

.PHONY: generate
generate: proto
