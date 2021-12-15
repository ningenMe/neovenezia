package service

import (
	"github.com/ningenme/neovenezia/pkg/model"
	"io/ioutil"
	"path/filepath"
)

func ExecTree(args []string) {
	fileNode := GetFileNode(getCurrentPath(args), "./", 0)
	fileNode.Print()
}

func getCurrentPath(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return "./"
}

func GetFileNode(path string, directoryName string, depth int) model.FileNode {

	files, err := ioutil.ReadDir(path)

	if err != nil {
		panic(err)
	}

	fileNode := model.FileNode{
		Path:          path,
		DirectoryName: directoryName,
		Files:         []string{},
		FileNodes:     []model.FileNode{},
		Depth:         depth,
	}

	for _, file := range files {
		fileName := file.Name()

		nextPath := filepath.Join(path, fileName)
		if file.IsDir() {
			childFileNode := GetFileNode(nextPath, fileName, depth+1)
			fileNode.FileNodes = append(fileNode.FileNodes, childFileNode)
		} else {
			fileNode.Files = append(fileNode.Files, fileName)
		}
	}
	return fileNode
}
