package templates

import (
	_ "embed"
)

//go:embed .golangci.yaml
var GolangCI []byte

//go:embed .gitignore
var Gitignore []byte

//go:embed buf.gen.server.yaml
var BufGenServer []byte

//go:embed buf.gen.client.yaml
var BufGenClient []byte

//go:embed buf.yaml
var Buf []byte

//go:embed .pre-commit.yaml
var PreCommit []byte

//go:embed Makefile
var Makefile []byte
