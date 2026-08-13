//go:build !prod

package main

import (
	"io/fs"
	"os"
)

var Assets fs.FS = os.DirFS("../..")
