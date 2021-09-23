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

	Op             int               `json:"op"`
	ServerID       string            `json:"serverID,omitempty"`
	GroupID        int               `json:"groupID,omitempty"`
	UserID         uint64            `json:"userID,omitempty"`
	Msg            []byte            `json:"msg"`
	Count          int               `json:"count"`
	GroupUserInfos map[string]string `json:"groupUserInfos"`
}

type TaskGroupPushParam struct {
	*DefaultTaskParam

	Op             int               `json:"op"`
	ServerID       string            `json:"serverID,omitempty"`
	GroupID        int               `json:"groupID,omitempty"`
	UserID         uint64            `json:"userID,omitempty"`
	Msg            []byte            `json:"msg"`
	Count          int               `json:"count"`
	GroupUserInfos map[string]string `json:"groupUserInfos"`
}

type TaskGroupCountParam struct {
	*DefaultTaskParam

	Op             int               `json:"op"`
	ServerID       string            `json:"serverID,omitempty"`
	GroupID        int               `json:"groupID,omitempty"`
	UserID         uint64            `json:"userID,omitempty"`
	Msg            []byte            `json:"msg"`
	Count          int               `json:"count"`
	GroupUserInfos map[string]string `json:"groupUserInfos"`
}

type TaskGroupInfoParam struct {
	*DefaultTaskParam

	Op             int               `json:"op"`
	ServerID       string            `json:"serverID,omitempty"`
	GroupID        int               `json:"groupID,omitempty"`
	UserID         uint64            `json:"userID,omitempty"`
	Msg            []byte            `json:"msg"`
	Count          int               `json:"count"`
	GroupUserInfos map[string]string `json:"groupUserInfos"`
}
