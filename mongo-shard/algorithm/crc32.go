package algorithm

import "hash/crc32"

type crc32Algorithm struct{}

func (c *crc32Algorithm) hash(data []byte) int {
	return int(crc32.ChecksumIEEE(data))
}
