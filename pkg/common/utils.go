package common

import (
	"fmt"
	"strings"
	"time"
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

func CreateTimestamp() string {
	return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
}
