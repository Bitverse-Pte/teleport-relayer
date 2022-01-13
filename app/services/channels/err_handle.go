package channels

import (
	"strings"
)

func handleRecvPacketsError(err error) bool {
	return strings.Contains(err.Error(), "packet has been received") || strings.Contains(err.Error(), "acknowledge packet verification failed: commitment bytes are not equal:")
}

func isBifurcate(err error) bool {
	return strings.Contains(err.Error(), "header does not exist for hash") || strings.Contains(err.Error(), "can not find consensus state for hash")
}
