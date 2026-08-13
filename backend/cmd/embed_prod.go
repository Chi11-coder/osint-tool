//go:build prod

package main

import "embed"

//go:embed all:dist
var Assets embed.FS
