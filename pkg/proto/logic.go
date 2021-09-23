package proto

// for api layer and connect layer rpc call

type LogicSigninArg struct {
	*DefaultRPCArg

	UserName string
	Password string
}

type LogicSigninReply struct {
	*DefaultRPCReply

	Code      int
	AuthToken string
}

type LogicSignupArg struct {
	*DefaultRPCArg

	UserName string
	Password string
}

type LogicSignupReply struct {
	*DefaultRPCReply

	Code      int
	AuthToken string
}

type LogicSignoutArg struct {
	*DefaultRPCArg

	AuthToken string
}

type LogicSignoutReply struct {
	*DefaultRPCReply

	Code int
}

type LogicAuthCheckArg struct {
	*DefaultRPCArg

	AuthToken string
}

type LogicAuthCheckReply struct {
	*DefaultRPCReply

	Code     int
	UserID   uint64
	UserName string
}

type LogicUserInfoQueryArg struct {
	*DefaultRPCArg

	UserID uint64
}

type LogicUserInfoQueryReply struct {
	*DefaultRPCReply

	Code     int
	UserID   uint64
	UserName string
}

type LogicPeerPushArg struct {
	*DefaultRPCArg

	Msg          string `json:"msg"`
	FromUserId   uint64 `json:"fromUserID"`
	FromUserName string `json:"fromUserName"`
	ToUserId     uint64 `json:"toUserID"`
	ToUserName   string `json:"toUserName"`
	GroupId      int    `json:"groupID"`
	Op           int    `json:"op"`
	Timestamp    string `json:"timestamp"`
}

type LogicPeerPushReply struct {
	*DefaultRPCReply

	Code int
	Msg  string
}

type LogicGroupPushArg struct {
	*DefaultRPCArg

	Msg          string `json:"msg"`
	FromUserId   uint64 `json:"fromUserID"`
	FromUserName string `json:"fromUserName"`
	ToUserId     uint64 `json:"toUserID"`
	ToUserName   string `json:"toUserName"`
	GroupId      int    `json:"groupID"`
	Op           int    `json:"op"`
	Timestamp    string `json:"timestamp"`
}

type LogicGroupPushReply struct {
	*DefaultRPCReply

	Code int
	Msg  string
}

type LogicGroupCountArg struct {
	*DefaultRPCArg

	Msg          string `json:"msg"`
	FromUserId   uint64 `json:"fromUserID"`
	FromUserName string `json:"fromUserName"`
	ToUserId     uint64 `json:"toUserID"`
	ToUserName   string `json:"toUserName"`
	GroupId      int    `json:"groupID"`
	Op           int    `json:"op"`
	Timestamp    string `json:"timestamp"`
}

type LogicGroupCountReply struct {
	*DefaultRPCReply

	Code int
	Msg  string
}

type LogicGroupInfoArg struct {
	*DefaultRPCArg

	Msg          string `json:"msg"`
	FromUserId   uint64 `json:"fromUserID"`
	FromUserName string `json:"fromUserName"`
	ToUserId     uint64 `json:"toUserID"`
	ToUserName   string `json:"toUserName"`
	GroupId      int    `json:"groupID"`
	Op           int    `json:"op"`
	Timestamp    string `json:"timestamp"`
}

type LogicGroupInfoReply struct {
	*DefaultRPCReply

	Code int
	Msg  string
}

type LogicConnectArg struct {
	*DefaultRPCArg

	AuthToken string `json:"authToken"`
	GroupID   int    `json:"groupID"`
	ServerID  int    `json:"serverID"`
}

type LogicConnectReply struct {
	*DefaultRPCReply

	Code   int
	UserID uint64
}

type LogicDisconnectArg struct {
	*DefaultRPCArg

	GroupID int
	UserID  uint64
}

type LogicDisconnectReply struct {
	*DefaultRPCReply

	Code int
}
