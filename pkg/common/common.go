package common

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
)

type Any = interface{}

const (
	NetworkSplitSign   = "@"
	ServerIDxSplitSign = "-"
)

func GetSnowflakeID(nodeID int64) string {
	node, _ := snowflake.NewNode(nodeID)
	return node.Generate().String()
}

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
	timestamp := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	return timestamp
}

func RandInt(mod int) int {
	rint := rand.Int() % mod
	return rint
}

func NewServerIDx(serverPath string, idx int) string {
	serverIDx := fmt.Sprintf("%s-%d", serverPath, idx)
	return serverIDx
}

// opt replace rand with slb rules
func GetServerIDx(serverPath string, idx int) string {
	serverIDx := fmt.Sprintf("%s-%d", serverPath, idx)
	return serverIDx
}
