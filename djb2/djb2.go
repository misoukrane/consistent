package djb2

// Hash - This is a straight port of HAPROXY's djb2 hash, because this is what we care about.
// We could of course make this more general by making the initial hash with a seed
// other than 5381
func Hash(s string) uint64 {
	var hash uint64
	hash = 5381
	length := len(s)
	var c uint64
	var i int

	for i = 0; length >= 8; length -= 8 {
		c = uint64(s[i])
		hash = ((hash << 5) + hash) + c
		i += 1

		c = uint64(s[i])
		hash = ((hash << 5) + hash) + c
		i += 1

		c = uint64(s[i])
		hash = ((hash << 5) + hash) + c
		i += 1

		c = uint64(s[i])
		hash = ((hash << 5) + hash) + c
		i += 1

		c = uint64(s[i])
		hash = ((hash << 5) + hash) + c
		i += 1

		c = uint64(s[i])
		hash = ((hash << 5) + hash) + c
		i += 1

		c = uint64(s[i])
		hash = ((hash << 5) + hash) + c
		i += 1

		c = uint64(s[i])
		hash = ((hash << 5) + hash) + c
		i += 1
	}

	switch length {
	case 7:
		c = uint64(s[i])
		hash = ((hash << 5) + hash) + c
		i += 1
		fallthrough
	case 6:
		c = uint64(s[i])
		hash = ((hash << 5) + hash) + c
		i += 1
		fallthrough
	case 5:
		c = uint64(s[i])
		hash = ((hash << 5) + hash) + c
		i += 1
		fallthrough
	case 4:
		c = uint64(s[i])
		hash = ((hash << 5) + hash) + c
		i += 1
		fallthrough
	case 3:
		c = uint64(s[i])
		hash = ((hash << 5) + hash) + c
		i += 1
		fallthrough
	case 2:
		c = uint64(s[i])
		hash = ((hash << 5) + hash) + c
		i += 1
		fallthrough
	case 1:
		c = uint64(s[i])
		hash = ((hash << 5) + hash) + c
		i += 1
		break
	default:
		break
	}

	return hash

}
