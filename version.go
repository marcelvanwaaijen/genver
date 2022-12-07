// Generated by genver.exe
// Created: 2022-12-07 12:27:20 CET
package main

import (
	"fmt"
	"os"
)

const (
    version string = "v1.0.2"
    gitcommit string = "v1.0.2"
)

func ShowVersionOnly() {
    fmt.Fprintf(os.Stdout, "version: %s\n", version)
    os.Exit(0)
}

func ShowVersion() {
    fmt.Fprintf(os.Stdout, "version: %s (%s)\n", version, gitcommit)
    os.Exit(0)
}