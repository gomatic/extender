package extension

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gomatic/commander"
)

//
func Delegate(args ...string) error {
	goroot, exists := os.LookupEnv("GOROOT")
	if !exists {
		return fmt.Errorf("Missing GOROOT")
	}
	cmd := commander.New("").Inherit(2)
	cmd.Binary = filepath.Join(goroot, "bin", "go")
	cmd.Args(args...)
	return cmd.Execute()
}
