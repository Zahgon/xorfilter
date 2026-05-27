//go:build (!amd64 && !386 && !arm && !arm64 && !ppc64le && !mipsle && !mips64le && !mips64p32le && !wasm) || appengine
// +build !amd64,!386,!arm,!arm64,!ppc64le,!mipsle,!mips64le,!mips64p32le,!wasm appengine

package xorfilter

import (
	"io"
)

// Save writes the filter to the writer in little endian format.
func (f *BinaryFuse[T]) Save(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Write the length of Fingerprints

// Write the Fingerprints

// LoadBinaryFuse reads the filter from the reader in little endian format.
func LoadBinaryFuse[T Unsigned](r io.Reader) (*BinaryFuse[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Read the length of Fingerprints
