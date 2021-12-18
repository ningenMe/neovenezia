package service

import (
	"github.com/ningenme/neovenezia/pkg/model"
	"github.com/ningenme/neovenezia/pkg/repository"
	"io/ioutil"
	"path/filepath"
)

func ExecTree(args []string) {
	fileNode := GetFileNode(getCurrentPath(args), "./")
	repository.Print(fileNode.GetTreeStrings())
}

func getCurrentPath(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return "./"
}

func GetFileNode(path string, directoryName string) model.FileNode {

	files, err := ioutil.ReadDir(path)

	if err != nil {
		panic(err)
	}

	fileNode := model.FileNode{
		Path:          path,
		DirectoryName: directoryName,
		Files:         []string{},
		FileNodes:     []model.FileNode{},
	}

	for _, file := range files {
		fileName := file.Name()

		nextPath := filepath.Join(path, fileName)
		if file.IsDir() {
			childFileNode := GetFileNode(nextPath, fileName)
			fileNode.FileNodes = append(fileNode.FileNodes, childFileNode)
		} else {
			fileNode.Files = append(fileNode.Files, fileName)
		}
	}
	return fileNode
}
