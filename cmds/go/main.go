package main

import (
	"log"
	"os"

	"path/filepath"

	"strings"

	"github.com/gomatic/commander"
	"github.com/gomatic/go-vbuild"
)

//
func main() {

	// Configure logging. TODO use a real logger.

	debugging, exists := os.LookupEnv("DEBUGGING")
	msgDebug := exists && strings.ToLower(debugging) == "true"

	log.Printf("Go toolchain extender v%s", build.Version.String())

	// Slice the prefixes.

	prefix, exists := os.LookupEnv("EXTENDER_PREFIX")
	if !exists {
		prefix = strings.Join([]string{"go-", "go"}, string(filepath.ListSeparator))
	}
	prefixes := filepath.SplitList(prefix)

	// Find the executable.

	var cmd *commander.Commanding
	for _, prefix := range prefixes {
		subcommand := os.Args[1]
		if msgDebug {
			log.Printf("Trying %s%s", prefix, subcommand)
		}
		c, err := commander.New(prefix).LookPath(subcommand)
		if err == nil { // Found
			// Inherit the environment and command-line starting after the subcommand.
			cmd = c.Inherit(1)
			break
		}
		if msgDebug {
			log.Println(err)
		}
	}

	// If an extension wasn't found, pass it on to GOROOT/bin/go

	if cmd == nil { // Not found
		if msgDebug {
			log.Printf("No extension found for %s", os.Args[1])
		}
		cmd = commander.New("").Inherit(0)
		goroot, exists := os.LookupEnv("GOROOT")
		if !exists {
			log.Println("Missing GOROOT")
			os.Exit(1)
		}
		cmd.Binary = filepath.Join(goroot, "bin", "go")
	}

	if msgDebug {
		log.Println(cmd.String())
	}

	//

	err := cmd.Execute()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
