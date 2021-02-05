package grains

import (
	"errors"
)

const MaxUint64 = (1 << 64) - 1 // << = x × 2^y

func Total () uint64 {
	return uint64(MaxUint64)
}

func Square (input int) (expected uint64, err error) {
	if 1 > input || input > 64 {
		return uint64(0), errors.New("invalid")
	}

	return uint64(1 << uint(input-1)), nil
}