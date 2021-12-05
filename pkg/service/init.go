package service

import (
	"fmt"
	"github.com/ningenme/neovenezia/pkg/repository"
	"os"
)

var neoveneziaYamlName string = "neovenezia.yaml"

func ExecInit() {

	if repository.FileExists(getNeoveneziaYamlFullPath()) {
		repository.PrintOne(getNeoveneziaYamlFullPath() + " already exists")
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
	repository.PrintOne("Initialized empty " + configFileFullPath)
}