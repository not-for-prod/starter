package internal

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/not-for-prod/starter/cmd/starter/internal/pkg/filewriter"
	"github.com/not-for-prod/starter/cmd/starter/internal/pkg/logger"
)

func overwrite(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, fs.ErrNotExist) {
		return true
	}

	keep := "don't overwrite " + path
	write := "overwrite " + path

	prompt := promptui.Select{
		Label: fmt.Sprintf("`%s` already exists, do you want to keep this file", filepath.Base(path)),
		Items: []string{keep, write},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return false
	}

	if result == write {
		return true
	}

	return false
}

func Save(path string, data []byte) {
	if !overwrite(path) {
		return
	}

	err := filewriter.WriteBytesToFile(path, data)
	if err != nil {
		logger.Fatalf(err.Error())
	}
}

func Execute(name string, arg ...string) {
	c := exec.Command(name, arg...)
	logger.Info("exec", c.String())

	err := c.Run()
	if err != nil {
		logger.Fatalf(err.Error())
	}
}
