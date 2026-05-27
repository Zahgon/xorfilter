package xorfilter

import "io"

type BinaryFuse8 BinaryFuse[uint8]

// PopulateBinaryFuse8 fills the filter with provided keys. For best results,
// the caller should avoid having too many duplicated keys.
// The function may return an error if the set is empty.
func PopulateBinaryFuse8(keys []uint64) (*BinaryFuse8, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Contains returns `true` if key is part of the set with a false positive probability of <0.4%.
func (filter *BinaryFuse8) Contains(key uint64) bool { _ = "STUB: not implemented"; return false }

// Save writes the filter to the writer in little endian format.
func (f *BinaryFuse8) Save(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// LoadBinaryFuse8 reads the filter from the reader in little endian format.
func LoadBinaryFuse8(r io.Reader) (*BinaryFuse8, error) { _ = "STUB: not implemented"; return nil, nil }
