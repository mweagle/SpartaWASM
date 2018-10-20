// +build mage

package main

import (
	"github.com/magefile/mage/sh"
	spartaMage "github.com/mweagle/Sparta/magefile"
)

// Build the WASM file
func BuildWASM() error {
	return sh.RunWith(map[string]string{
		"GOARCH": "wasm",
		"GOOS":   "js",
	}, "go",
		"build",
		"-o",
		"./resources/sparta.wasm",
		"./wasm/main.go")
}

// Provision the service
func Provision() error {
	return spartaMage.Provision()
}

// Describe the stack by producing an HTML representation of the CloudFormation
// template
func Describe() error {
	return spartaMage.Describe()
}

// Delete the service, iff it exists
func Delete() error {
	return spartaMage.Delete()
}

// Status report if the stack has been provisioned
func Status() error {
	return spartaMage.Status()
}

// Version information
func Version() error {
	return spartaMage.Version()
}
