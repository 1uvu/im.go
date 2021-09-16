package proto

const (
	opUnknown = -100

	OpPeerPush = iota
	OpGroupPush
	OpGroupCount
	OpGroupInfo
	OpBuildTCPConn
)

var opTexts = map[int]string{
	opUnknown: "Unknown",

	OpPeerPush:     "PeerPush",
	OpGroupPush:    "GroupPush",
	OpGroupCount:   "GroupCount",
	OpGroupInfo:    "GroupInfo",
	OpBuildTCPConn: "BuildTCPConn",
}

func OPText(op int) string {
	if opText, ok := opTexts[op]; ok {
		return opText
	}

	return opTexts[opUnknown]
}
