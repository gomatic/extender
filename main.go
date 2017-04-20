package main

import (
	"fmt"
	"log"
)

//
const majorVersion = "1.0"

//
var appVersion = "0"

//
func main() {
	log.Printf("Go toolchain extender v%s.%s", majorVersion, appVersion)

	// ensure extender's path comes before go in PATH

	extender := "${GOBIN}"

	// update PATH for this process

	// and check that `go` is ours
	// strings.HasPrefix(os.StartProcess("go version"), "Go toolchain extender")

	// output for eval
	fmt.Printf("export PATH=%s:${PATH}\n", extender)
}
