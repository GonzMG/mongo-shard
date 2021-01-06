package algorithm

import "github.com/howeyc/crc16"

type crc16Algorithm struct{}

func (c *crc16Algorithm) hash(data []byte) int {
	return int(crc16.ChecksumCCITTFalse(data))
}
