package version

import (
	"github.com/ningenme/neovenetia/pkg/common"
)

// if you update version, please update CHANGELOG too.
func getVersion() string {
	return "v0.2.0"
}

func Exec() {
	var message string = "Neovenetia " + getVersion()
	common.PrintOne(message)
}