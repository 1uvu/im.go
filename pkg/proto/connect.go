package proto

type Msg struct {
	Ver       int    `json:"ver"`
	Operation int    `json:"op"`
	SeqID     string `json:"seq"`
	Body      []byte `json:"body"`
}

type PeerPushRequest struct {
	UserID uint64
	Msg    Msg
}

type GroupPushRequest struct {
	GroupID int
	Msg     Msg
}

type GroupCountRequest struct {
	GroupID int
	Count   uint64
}
