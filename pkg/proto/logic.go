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
	UserID   int
	UserName string
}

type UserInfoQueryArg struct {
	*LogicArg

	UserID int
}

type UserInfoQueryReply struct {
	*LogicReply

	Code     int
	UserID   int
	UserName string
}

type OpArg struct {
	*LogicArg

	Msg          string
	FromUserId   int
	FromUserName string
	ToUserId     int
	ToUserName   string
	GroupId      int
	Op           int
}

type OpReply struct {
	*LogicReply

	Code int
	Msg  string
}

type ConnectArg struct {
	*LogicArg

	AuthToken string `json:"authToken"`
	GroupID   uint64 `json:"groupID"`
	ServerID  int    `json:"serverID"`
}

type ConnectReply struct {
	*LogicReply

	Code   int
	UserID uint64
}

type DisconnectArg struct {
	*LogicArg

	GroupID uint64
	UserID  uint64
}

type DisconnectReply struct {
	*LogicReply

	Code int
}
