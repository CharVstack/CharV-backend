package util

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

func SearchVmId(socks []string, x openapi_types.UUID) int {
	for i, v := range socks {
		if v == x.String() {
			return i
		}
	}
	return -1
}
