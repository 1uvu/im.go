package proto

// for task layer rpc call

type Msg struct {
	*DefaultRPCArg

	Ver       int    `json:"ver"`
	Operation int    `json:"op"`
	SeqID     string `json:"seq"`
	Body      []byte `json:"body"`
}

type ConnectPeerArg struct {
	*DefaultRPCArg

	UserID uint64
	Msg    Msg
}

type ConnectGroupArg struct {
	*DefaultRPCArg

	GroupID int
	Msg     Msg
}

type ConnectReply struct {
	*DefaultRPCReply

	Code int
}
