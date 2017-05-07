package extension

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gomatic/commander"
)

//
func Delegate(subcommand string, args ...string) error {
	goroot, exists := os.LookupEnv("GOROOT")
	if !exists {
		return fmt.Errorf("Missing GOROOT")
	}
	cmd := commander.New("")
	cmd.Binary = filepath.Join(goroot, "bin", "go")
	cmd.Args(subcommand).Args(args...).Inherit(0)
	return cmd.Execute()
}
