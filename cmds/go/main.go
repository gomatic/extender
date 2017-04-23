package main

import (
	"log"
	"os"

	"path/filepath"

	"strings"

	"github.com/gomatic/commander"
)

//
const majorVersion = "1.0"

//
var appVersion = "0"

//
func main() {

	// Configure logging. TODO use a real logger.

	verbosity, exists := os.LookupEnv("EXTENDER_MESSAGING")
	verbosity = strings.ToLower(verbosity)
	msgSilent := verbosity == "silent"
	msgDebug := !msgSilent && verbosity == "debug"
	msgInfo := !msgSilent && msgDebug || verbosity == "info"

	if msgInfo {
		log.Printf("Go toolchain extender v%s.%s", majorVersion, appVersion)
	}

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
			cmd = c.Inherit(2)
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
		cmd = commander.New("").Inherit(1)
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
