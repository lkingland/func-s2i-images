package main

import (
	"fmt"
	"os"
	"runtime/debug"
)

var usage = `func-s2i-images

Build and publish S2I-ready images for use by Functions builtin langauges.

Example:
  func-s2i-images build && func-s2i-images publish

Commands:
  build:    Build images locally
	publish:  Publish images publicly
	version:  Print version information
`

func main() {
	fmt.Println("Hello World")
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Fprintf(os.Stderr, "unable to read build information")
	}
	fmt.Println(bi.Main.Version)
}
