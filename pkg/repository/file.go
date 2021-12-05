package repository

import (
	"github.com/ningenme/neovenezia/pkg/model"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func FileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func GetFiles(directory string) []string {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		fileName := file.Name()

		if model.IsInclude(getExcludeDirectoryName(), fileName) {
			continue
		}

		if isFirstDot(fileName) {
			continue
		}

		var fileFullPath = filepath.Join(directory, fileName)
		if file.IsDir() {
			paths = append(paths, GetFiles(fileFullPath)...)
			continue
		}
		paths = append(paths, fileFullPath)
	}
	return paths
}

func getExcludeDirectoryName() []string {
	return []string{".git", ".idea"}
}

func isFirstDot(fileName string) bool {
	return regexp.MustCompile(`^\..*`).Match([]byte(fileName))
}
