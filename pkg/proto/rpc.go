package proto

type IRPCArg interface {
	MustEmbedDefaultRPCArg()
}

type IRPCReply interface {
	GetErrMsg() string
	SetErrMsg(string)
	MustEmbedDefaultRPCReply()
}

type DefaultRPCArg struct{}

func (arg *DefaultRPCArg) MustEmbedDefaultRPCArg() {
}

type DefaultRPCReply struct {
	ErrMsg string `json:"errMsg"`
}

func (reply *DefaultRPCReply) GetErrMsg() string {
	return reply.ErrMsg
}

func (reply *DefaultRPCReply) SetErrMsg(msg string) {
	reply.ErrMsg = msg
}

func (reply *DefaultRPCReply) MustEmbedDefaultRPCReply() {

}
