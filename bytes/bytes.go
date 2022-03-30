package bytes

// Reverse takes a slice of bytes and returns a reverse order byte slice.
func Reverse(bb []byte) []byte {
	br := []byte{}
	for i := len(bb) - 1; i >= 0; i-- {
		br = append(br, bb[i])
	}
	return br
}
