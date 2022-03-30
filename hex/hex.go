package hex

import "strconv"

func HexToUint64(h string) (uint64, error) {
	i, err := strconv.ParseInt(h, 16, 64)
	return uint64(i), err
}
