package model

import (
	"fmt"
)

type FileNode struct {
	Directory string
	FileName  string
	Files     []string
	FileNodes []FileNode
}

func (fileNode *FileNode) GetDirectories() []string {
	var directories []string
	for _, childrenFileNode := range fileNode.FileNodes {
		directories = append(directories, childrenFileNode.FileName)
	}
	return directories
}

func (fileNode *FileNode) Print() {
	fmt.Println(fileNode.Directory)
	for idx, file := range fileNode.Files {

		prefix := "├──"
		if idx+1 == len(fileNode.Files) {
			prefix = "└──"
		}
		fmt.Println(prefix, file)
	}
}
