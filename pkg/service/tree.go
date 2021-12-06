package service

import (
	"fmt"
	"github.com/ningenme/neovenezia/pkg/model"
	"io/ioutil"
	"path/filepath"
)

func ExecTree() {
	//repository.Print(GetFiles("./"))
	fileNode := GetFileNode("./")
	fmt.Println(fileNode.FileName)
	fmt.Println(fileNode.Files)
	fmt.Println(fileNode.GetDirectories())
	fileNode.Print()
}

func GetFiles(directory string) []string {
	files, err := ioutil.ReadDir(directory)

	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		fileName := file.Name()

		var fileFullPath = filepath.Join(directory, fileName)
		if file.IsDir() {
			paths = append(paths, GetFiles(fileFullPath)...)
			continue
		}
		paths = append(paths, fileFullPath)
	}
	return paths
}

func GetFileNode(directory string) model.FileNode {
	files, err := ioutil.ReadDir(directory)

	if err != nil {
		panic(err)
	}

	fileNode := model.FileNode{Directory: directory, FileName: directory, Files: []string{}, FileNodes: []model.FileNode{}}
	fileNode.Directory = directory

	for _, file := range files {
		fileName := file.Name()

		var fileFullPath = filepath.Join(directory, fileName)
		if file.IsDir() {
			childFileNode := GetFileNode(fileFullPath)
			fileNode.FileNodes = append(fileNode.FileNodes, childFileNode)
		} else {
			fileNode.Files = append(fileNode.Files, fileFullPath)
		}
	}
	return fileNode
}
