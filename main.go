package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//
const majorVersion = "1.0"

//
var appVersion = "0"

//
func main() {
	log.Printf("Go toolchain extender v%s.%s", majorVersion, appVersion)

	// Ensure extender's path, GOBIN, comes before go in PATH.

	// If there's no GOBIN, binaries will be placed into GOPATH[0]/bin
	gobin, exists := os.LookupEnv("GOBIN")
	if !exists || gobin == "" {
		gopath, exists := os.LookupEnv("GOPATH")
		if !exists || gopath == "" {
			log.Println("Missing GOBIN and GOPATH")
			os.Exit(1)
		}
		gobin = filepath.Join(filepath.SplitList(gopath)[0], "bin")
	}
	// TODO If there's not a GOROOT, use the go executable's path as long as it's not GOBIN.
	goroot, exists := os.LookupEnv("GOROOT")
	if !exists || goroot == "" {
		log.Println("Missing GOROOT")
		os.Exit(1)
	}
	//
	path, exists := os.LookupEnv("PATH")
	if !exists {
		log.Println("Missing PATH")
		os.Exit(1)
	}

	// Use absolute paths for verification.
	gobinAbs, err := filepath.Abs(gobin)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	gorootAbs, err := filepath.Abs(goroot)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// TODO support the current shell+OS

	if args := os.Args[1:]; len(args) > 0 {
		fmt.Printf("export EXTENDER_PREFIX=%s\n", strings.Join(args, string(filepath.ListSeparator)))
	}

	// Ensure that the location of go get binaries precedes GOROOT/bin.

	for _, path := range filepath.SplitList(path) {
		pathAbs, err := filepath.Abs(path)
		if err != nil {
			log.Println(err)
			continue
		}
		if strings.HasPrefix(pathAbs, gobinAbs) {
			// If we reach GOBIN first, then PATH is good.
			os.Exit(0)
		} else if strings.HasPrefix(pathAbs, gorootAbs) {
			// If GOROOT is before GOBIN or neither is in the PATH (unlikely), update the PATH.
			break
		}
	}

	fmt.Printf("export PATH=%s:${PATH}\n", gobin)
}
