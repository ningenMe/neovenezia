package version

import "fmt"

// if you update version, please update CHANGELOG too.
var version string = "v0.2.0"

func Main() {
	fmt.Println("Neovenetia " + version)
}