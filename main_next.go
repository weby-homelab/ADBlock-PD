//go:build next

package main

import (
	"embed"

	"github.com/weby-homelab/adblock-pd/internal/next/cmd"
	"github.com/weby-homelab/adblock-pd/internal/version"
)

// Embed the prebuilt client here since we strive to keep .go files inside the
// internal directory and the embed package is unable to embed files located
// outside of the same or underlying directory.

//go:embed build
var frontend embed.FS

//go:embed client/package.json
var clientPackageJSON []byte

func main() {
	version.SetFromPackageJSON(clientPackageJSON)
	cmd.Main(frontend)
}
