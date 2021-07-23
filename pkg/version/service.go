package version

import (
	"github.com/ningenme/neovenezia/pkg/common"
)

// if you update version, please update CHANGELOG too.
func getVersion() string {
	return "v0.3.0"
}

func Exec() {
	var message string = "Neovenezia " + getVersion()
	common.PrintOne(message)
}