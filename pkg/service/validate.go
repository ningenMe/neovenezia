package service

const defaultDirectory = "."
func ExecValidate(args []string) {
	//var directory = getDirectory(args)
	//var paths = repository.GetFiles(directory)
	//paths = getPaths(paths, directory)
	//repository.Print(paths)
}

func getDirectory(args []string) string {
	if len(args) >= 1 {
		return args[0]
	}
	return defaultDirectory
}

func getPaths(paths []string, directory string) []string  {
	if directory == defaultDirectory {
		return paths
	}
	var trimSize = len(directory)

	var newPaths []string
	for _, path := range paths {
		newPaths = append(newPaths, path[trimSize:])
	}
	return newPaths
}
