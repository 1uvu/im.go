package common

import (
	"fmt"
	"strings"
)

const (
	NetworkSplitSign = "@"
)

func ParseNetworkAddr(str string) (network, addr string, err error) {
	if idx := strings.Index(str, NetworkSplitSign); idx == -1 {
		err = fmt.Errorf("addr: %s format got error, which is must be network@tcp:port or network@unixsocket", str)
	} else {
		network = str[:idx]
		addr = str[idx+1:]
	}
	return
}
