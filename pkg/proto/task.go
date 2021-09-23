package proto

// for logic layer rpc call

type ITaskParam interface {
	GetOp() int
	MustEmbedDefaultTaskParam()
}

type DefaultTaskParam struct {
	Op int
}

func (param *DefaultTaskParam) MustEmbedDefaultTaskParam() {
}

func (param *DefaultTaskParam) GetOp() int {
	return param.Op
}

type TaskPeerPushParam struct {
	*DefaultTaskParam

	Op        int    `json:"op"`
	ServerIDx string `json:"serverIDx,omitempty"`
	UserID    uint64 `json:"userID,omitempty"`
	Msg       []byte `json:"msg"`
}

type TaskGroupPushParam struct {
	*DefaultTaskParam

	Op      int    `json:"op"`
	GroupID int    `json:"groupID,omitempty"`
	Msg     []byte `json:"msg"`
}

type TaskGroupCountParam struct {
	*DefaultTaskParam

	Op      int    `json:"op"`
	GroupID int    `json:"groupID,omitempty"`
	Count   uint64 `json:"count"`
}

type TaskGroupInfoParam struct {
	*DefaultTaskParam

	Op             int               `json:"op"`
	GroupID        int               `json:"groupID,omitempty"`
	GroupUserInfos map[string]string `json:"groupUserInfos"`
}
