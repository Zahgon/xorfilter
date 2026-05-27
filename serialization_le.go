//go:build amd64 || 386 || arm || arm64 || ppc64le || mipsle || mips64le || mips64p32le || wasm

package xorfilter

import (
	"io"
)

// Save writes the filter to the writer assuming little endian system, using direct byte copy for performance.
func (f *BinaryFuse[T]) Save(w io.Writer) error {
	_ = "STUB: not implemented"
	// Write Seed
	return nil
}

// Write SegmentLength

// Write SegmentLengthMask

// Write SegmentCount

// Write SegmentCountLength

// Write length of Fingerprints

// Write Fingerprints

// LoadBinaryFuse reads the filter from the reader assuming little endian system, using direct byte copy for performance.
func LoadBinaryFuse[T Unsigned](r io.Reader) (*BinaryFuse[T], error) {
	_ = "STUB: not implemented"

	// Read Seed
	return nil, nil
}

// Read SegmentLength

// Read SegmentLengthMask

// Read SegmentCount

// Read SegmentCountLength

// Read length of Fingerprints
