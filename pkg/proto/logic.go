package proto

type ILogicArg interface {
	MustEmbedDefaultLogicArg()
}

type ILogicReply interface {
	GetErrMsg() string
	SetErrMsg(string)
	MustEmbedDefaultLogicReply()
}

type LogicArg struct{}

func (arg *LogicArg) MustEmbedDefaultLogicArg() {
}

type LogicReply struct {
	ErrMsg string `json:"errMsg"`
}

func (reply *LogicReply) GetErrMsg() string {
	return reply.ErrMsg
}

func (reply *LogicReply) SetErrMsg(msg string) {
	reply.ErrMsg = msg
}

func (reply *LogicReply) MustEmbedDefaultLogicReply() {

}

type SigninArg struct {
	*LogicArg

	UserName string
	Password string
}

type SigninReply struct {
	*LogicReply

	Code      int
	AuthToken string
}

type SignupArg struct {
	*LogicArg

	UserName string
	Password string
}

type SignupReply struct {
	*LogicReply

	Code      int
	AuthToken string
}

type SignoutArg struct {
	*LogicArg

	AuthToken string
}

type SignoutReply struct {
	*LogicReply

	Code int
}

type AuthCheckArg struct {
	*LogicArg

	AuthToken string
}

type AuthCheckReply struct {
	*LogicReply

	Code     int
	UserID   uint64
	UserName string
}

type UserInfoQueryArg struct {
	*LogicArg

	UserID uint64
}

type UserInfoQueryReply struct {
	*LogicReply

	Code     int
	UserID   uint64
	UserName string
}

type PushArg struct {
	*LogicArg

	Msg          string `json:"msg"`
	FromUserId   uint64 `json:"fromUserID"`
	FromUserName string `json:"fromUserName"`
	ToUserId     uint64 `json:"toUserID"`
	ToUserName   string `json:"toUserName"`
	GroupId      int    `json:"groupID"`
	Op           int    `json:"op"`
	Timestamp    string `json:"timestamp"`
}

type PushReply struct {
	*LogicReply

	Code int
	Msg  string
}

type ConnectArg struct {
	*LogicArg

	AuthToken string `json:"authToken"`
	GroupID   int    `json:"groupID"`
	ServerID  int    `json:"serverID"`
}

type ConnectReply struct {
	*LogicReply

	Code   int
	UserID uint64
}

type DisconnectArg struct {
	*LogicArg

	GroupID int
	UserID  uint64
}

type DisconnectReply struct {
	*LogicReply

	Code int
}
