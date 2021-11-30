package service

import (
	"fmt"
	"github.com/ningenme/neovenezia/pkg/common"
	"os"
)

var neoveneziaYamlName string = "neovenezia.yaml"

func ExecInit() {

	if common.FileExists(getNeoveneziaYamlFullPath()) {
		common.PrintOne(getNeoveneziaYamlFullPath() + " already exists")
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