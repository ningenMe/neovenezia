package service

import (
	"github.com/ningenme/neovenezia/pkg/repository"
)

// if you update version, please update CHANGELOG too.
func getVersion() string {
	return "v0.3.0"
}

func ExecVersion() {
	var message string = "Neovenezia " + getVersion()
	repository.PrintOne(message)
}
