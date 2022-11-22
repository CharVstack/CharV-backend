package util

import (
	"github.com/google/uuid"
)

func SearchVmId(socks []string, x uuid.UUID) int {
	for i, v := range socks {
		if v == x.String() {
			return i
		}
	}
	return -1
}
