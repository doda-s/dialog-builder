package core

import (
	"os"
)

type Directories struct {
	Cwd string
}

var Directory Directories

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		cwd = ""
	}

	Directory = Directories{
		Cwd: cwd,
	}
}