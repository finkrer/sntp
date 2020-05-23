package main

import (
	"errors"
	"math/bits"
)

func setBit(b *byte, n int) {
	*b |= (1 << n)
}

func bitSet(b byte, n int) bool {
	return b&(1<<n) > 0
}

func readBits(b byte, start int, len int) uint8 {
	result := uint8(0)
	for i := 0; i < len; i++ {
		if bitSet(b, i) {
			result += 1 << i
		}
	}
	return result
}

func writeBits(b *byte, start int, len int, val uint8) error {
	if bits.Len(uint(val)) > (start+len) || start+len > 8 {
		return errors.New("Value was too big for the length specified")
	}
	bval := byte(val)
	for i := 0; i < len; i++ {
		if bitSet(bval, i) {
			setBit(b, i+(8-start-len))
		}
	}
	return nil
}
