package xorfilter

func murmur64(h uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// returns random number, modifies the seed
func splitmix64(seed *uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func mixsplit(key, seed uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func rotl64(n uint64, c int) uint64 { _ = "STUB: not implemented"; return 0 }

func reduce(hash, n uint32) uint32 {
	_ = "STUB: not implemented"
	// http://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
	return 0
}

func fingerprint(hash uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// Contains tell you whether the key is likely part of the set
func (filter *Xor8) Contains(key uint64) bool { _ = "STUB: not implemented"; return false }

func (filter *Xor8) geth0h1h2(k uint64) hashes { _ = "STUB: not implemented"; return *new(hashes) }

func (filter *Xor8) geth0(hash uint64) uint32 { _ = "STUB: not implemented"; return 0 }

func (filter *Xor8) geth1(hash uint64) uint32 { _ = "STUB: not implemented"; return 0 }

func (filter *Xor8) geth2(hash uint64) uint32 { _ = "STUB: not implemented"; return 0 }

// scan for values with a count of one
func scanCount(Qi []keyindex, setsi []xorset) ([]keyindex, int) {
	_ = "STUB: not implemented"

	// len(setsi) = filter.BlockLength
	return nil, 0
}

// MaxIterations is the maximum number of iterations allowed before the populate
// function returns an error.
var MaxIterations = 1024

// Populate fills the filter with provided keys. For best results,
// the caller should avoid having too many duplicated keys.
// The function may return an error if the set is empty.
func Populate(keys []uint64) (*Xor8, error) { _ = "STUB: not implemented"; return nil, nil }

// round it down to a multiple of 3

// slice capacity defaults to length

// The probability of this happening is lower than the
// the cosmic-ray probability (i.e., a cosmic ray corrupts your system).

// scan for values with a count of one

// not actually possible after the initial scan.

// success

func pruneDuplicates(array []uint64) []uint64 { _ = "STUB: not implemented"; return nil }
