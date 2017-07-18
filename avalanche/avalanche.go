package avalanche

// Implementation of haproxy's full_hash avalanche.
// Reference: https://github.com/haproxy/haproxy/blob/25f067ccec52f53b0248a05caceb7841a3cb99df/include/common/standard.h#L740

const (
	BJ_CONST uint64 = 3221225473
)

func Avalanche(i uint64) uint64 {
	i = (i + 0x7ed55d16) + (i << 12)
	i = (i ^ 0xc761c23c) ^ (i >> 19)
	i = (i + 0x165667b1) + (i << 5)
	i = (i + 0xd3a2646c) ^ (i << 9)
	i = (i + 0xfd7046c5) + (i << 3)
	i = (i ^ 0xb55a4f09) ^ (i >> 16)

	return i * BJ_CONST
}
