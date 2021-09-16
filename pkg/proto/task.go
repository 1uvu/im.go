package proto

// type ITaskArg interface {
// 	MustEmbedDefaultTaskArg()
// }

// type ITaskReply interface {
// 	GetErrMsg() string
// 	SetErrMsg(string)
// 	MustEmbedDefaultTaskReply()
// }

// type TaskArg struct{}

// func (arg *TaskArg) MustEmbedDefaultTaskArg() {
// }

// type TaskReply struct {
// 	ErrMsg string `json:"errMsg"`
// }

// func (reply *TaskReply) GetErrMsg() string {
// 	return reply.ErrMsg
// }

// func (reply *TaskReply) SetErrMsg(msg string) {
// 	reply.ErrMsg = msg
// }

// func (reply *TaskReply) MustEmbedDefaultTaskReply() {

// }

type PublishArg struct {
	Op             int               `json:"op"`
	ServerID       string            `json:"serverID,omitempty"`
	GroupID        int               `json:"groupID,omitempty"`
	UserID         uint64            `json:"userID,omitempty"`
	Msg            []byte            `json:"msg"`
	Count          int               `json:"count"`
	GroupUserInfos map[string]string `json:"groupUserInfos"`
}

type PublishGroupInfoArg struct {
	Op             int               `json:"op"`
	GroupID        int               `json:"groupID,omitempty"`
	Count          uint64            `json:"count,omitempty"`
	GroupUserInfos map[string]string `json:"groupUserInfos"`
}

type PublishGroupCountArg struct {
	Op    int    `json:"op"`
	Count uint64 `json:"count,omitempty"`
}

type PublishReply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
