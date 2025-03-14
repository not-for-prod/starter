package main

import (
	"github.com/not-for-prod/starter/cmd/start/internal"
	"github.com/not-for-prod/starter/cmd/start/internal/pkg/logger"
	"github.com/not-for-prod/starter/cmd/start/internal/templates"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "starter",
		Short: "starter initiates project",
		Long:  `creates basic files for project`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				logger.Fatalf("module name required")
			}

			module := args[0]
			logger.Info("module", module)

			// create all basic files
			internal.Save(".gitignore", templates.Gitignore)
			internal.Save(".golangci.yaml", templates.GolangCI)
			internal.Save(".pre-commit.yaml", templates.PreCommit)
			internal.Save("buf.gen.server.yaml", templates.BufGenServer)
			internal.Save("buf.gen.client.yaml", templates.BufGenClient)
			internal.Save("buf.yaml", templates.Buf)
			internal.Save("Makefile", templates.Makefile)

			// go mod init
			internal.Execute("go", "mod", "init", module)

			// go mod tidy
			internal.Execute("go", "mod", "tidy")
		},
	}

	_ = rootCmd.Execute()
}
