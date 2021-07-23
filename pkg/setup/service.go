package setup

import (
	"fmt"
	"github.com/ningenme/neovenezia/pkg/common"
	"os"
)

var neoveneziaYamlName string = "neovenezia.yaml"

func Exec() {

	if common.FileExists(getNeoveneziaYamlFullPath()) {
		common.PrintOne(getNeoveneziaYamlFullPath() + " exists")
		return
	}

	createNeoveneziaYaml(getNeoveneziaYamlFullPath())
}

func getNeoveneziaYamlFullPath() string {
	currentPath, _ := os.Getwd()
	return currentPath + "/" + neoveneziaYamlName
}

func createNeoveneziaYaml(configFileFullPath string) {
	fp, err := os.Create(configFileFullPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	common.PrintOne("Initialized empty " + configFileFullPath)
}