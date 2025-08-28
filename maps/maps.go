/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package maps implements the functions, types, and interfaces for the module.
package maps

import (
	_ "golang.org/x/exp/maps"
)

//go:generate adptool .
//go:adapter:package golang.org/x/exp/maps
