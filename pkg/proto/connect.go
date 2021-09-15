package proto

type Msg struct {
	Ver       int    `json:"ver"`
	Operation int    `json:"op"`
	SeqID     string `json:"seq"`
	Body      []byte `json:"body"`
}

type PeerChatRequest struct {
	UserID uint64
	Msg    Msg
}

type GroupChatRequest struct {
	GroupID uint64
	Msg     Msg
}

type GroupCountRequest struct {
	GroupID uint64
	Count   uint64
}
