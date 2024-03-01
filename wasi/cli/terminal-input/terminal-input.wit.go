// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package terminalinput represents the interface "wasi:cli/terminal-input@0.2.0".
//
// Terminal input.
//
// In the future, this may include functions for disabling echoing,
// disabling input buffering so that keyboard events are sent through
// immediately, querying supported features, and so on.
package terminalinput

import (
	"github.com/ydnar/wasm-tools-go/cm"
)

// TerminalInput represents the resource "wasi:cli/terminal-input@0.2.0#terminal-input".
//
// The input side of a terminal.
//
//	resource terminal-input
type TerminalInput cm.Resource

// ResourceDrop represents the Canonical ABI function "resource-drop".
//
// Drops a resource handle.
//
//go:nosplit
func (self TerminalInput) ResourceDrop() {
	self.resourceDrop()
}

//go:wasmimport wasi:cli/terminal-input@0.2.0 [resource-drop]terminal-input
//go:noescape
func (self TerminalInput) resourceDrop()
