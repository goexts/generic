package maps

import (
	// Blank import to ensure the `adptool` generator can resolve the package.
	_ "golang.org/x/exp/maps"
)

//go:generate adptool .
//go:adapter:package golang.org/x/exp/maps
