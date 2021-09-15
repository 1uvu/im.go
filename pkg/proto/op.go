package proto

const (
	opUnknown = -100

	OpPeerChat = iota
	OpGroupChat
	OpGroupCount
	OpGroupInfo
	OpBuildTCPConn
)

var opTexts = map[int]string{
	opUnknown: "Unknown",

	OpPeerChat:     "PeerChat",
	OpGroupChat:    "GroupChat",
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
