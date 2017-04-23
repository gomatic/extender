package main

import (
	"log"
)

//
const majorVersion = "1.0"

//
var appVersion = "0"

//
func main() {
	log.Printf("Example toolchain extender v%s.%s", majorVersion, appVersion)
}
