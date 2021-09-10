package proto

type Msg struct {
	Ver       int    `json:"ver"`
	Operation int    `json:"op"`
	SeqID     string `json:"seq"`
	Body      []byte `json:"body"`
}

type PeerChatArg struct {
	UserID uint64
	Msg    Msg
}

type GroupChatArg struct {
	GroupID uint64
	Msg     Msg
}

type GroupCountArg struct {
	GroupID uint64
	Count   uint64
}
