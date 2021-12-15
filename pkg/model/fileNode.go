package model

import (
	"fmt"
	"strings"
)

type FileNode struct {
	Path          string
	DirectoryName string
	Files         []string
	FileNodes     []FileNode
	Depth         int
}

const (
	straight     = "│  "
	normalBranch = "├──"
	lastBranch   = "└──"
	blank        = "   "
)

func (fileNode *FileNode) GetDirectories() []string {
	var directories []string
	for _, childrenFileNode := range fileNode.FileNodes {
		directories = append(directories, childrenFileNode.DirectoryName)
	}
	return directories
}

func getExcludeDirectoryName() []string {
	return []string{".git", ".idea"}
}

func (fileNode *FileNode) Print() {
	directoryName := fileNode.DirectoryName

	fmt.Println(getDirectoryName(fileNode.Depth, fileNode.Path, directoryName))

	if IsInclude(getExcludeDirectoryName(), directoryName) {
		return
	}

	length := len(fileNode.Files)
	for index, file := range fileNode.Files {
		prefix := getFilePrefix(fileNode.Depth, index, length)
		fmt.Println(prefix, file)
	}
	for _, childrenFileNode := range fileNode.FileNodes {
		childrenFileNode.Print()
	}
}

func getDirectoryName(depth int, path string, directory string) string {
	if depth == 0 {
		return path
	}
	return strings.Repeat(straight, depth-1) + getBranch(0, 2) + " " + directory
}

func getFilePrefix(depth int, index int, length int) string {
	return strings.Repeat(straight, depth) + getBranch(index, length)
}

func getBranch(index int, length int) string {
	if index+1 == length {
		return lastBranch
	}
	return normalBranch
}
