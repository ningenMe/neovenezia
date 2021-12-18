package service

import (
	"github.com/ningenme/neovenezia/pkg/repository"
)

// if you update version, please update CHANGELOG too.
func getVersion() string {
	return "v0.4.0"
}

func ExecVersion() {
	message := "Neovenezia " + getVersion()
	repository.PrintOne(message)
}
