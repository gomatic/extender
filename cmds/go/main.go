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
	verbosity, exists := os.LookupEnv("EXTENDER_MESSAGING")
	verbosity = strings.ToLower(verbosity)
	messages_silent := verbosity == "silent"
	messages_debug := !messages_silent && verbosity == "debug"
	messages_info := !messages_silent && messages_debug || verbosity == "info"

	if messages_info {
		log.Printf("Go toolchain extender v%s.%s", majorVersion, appVersion)
	}

	prefix, exists := os.LookupEnv("EXTENDER_PREFIX")
	if !exists {
		prefix = "go-"
	}
	prefixes := strings.Split(prefix, ":")

	var cmd *commander.Commanding
	for _, prefix := range prefixes {
		subcommand := os.Args[1]
		if messages_debug {
			log.Printf("Trying %s%s", prefix, subcommand)
		}
		c, err := commander.New(prefix).Inherit(1).LookPath(subcommand)
		if err != nil {
			if messages_debug {
				log.Println(err)
			}
			continue
		}
		c.Cmd.Args = c.Cmd.Args[1:]
		cmd = c
		break
	}

	if cmd == nil {
		if messages_debug {
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

	if messages_debug {
		log.Println(cmd.String())
	}

	err := cmd.Execute()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
