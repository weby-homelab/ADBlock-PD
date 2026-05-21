//go:build !next

package main

import (
	"embed"
	// Embed tzdata in binary.
	//
	// See https://github.com/weby-homelab/adblock-pd/issues/6758
	_ "time/tzdata"

	"github.com/weby-homelab/adblock-pd/internal/home"
	"github.com/weby-homelab/adblock-pd/internal/version"
)

// Embed the prebuilt client here since we strive to keep .go files inside the
// internal directory and the embed package is unable to embed files located
// outside of the same or underlying directory.

//go:embed build
var clientBuildFS embed.FS

//go:embed client/package.json
var clientPackageJSON []byte

func main() {
	version.SetFromPackageJSON(clientPackageJSON)
	home.Main(clientBuildFS)
}
