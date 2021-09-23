package proto

import "im/pkg/common"

type APIPeerPushRequest struct {
	Msg       string `form:"msg" json:"msg" binding:"required"`
	ToUserID  string `form:"toUserID" json:"toUserID" binding:"required"`
	ToGroupID string `form:"toGroupID" json:"toGroupID" binding:"required"`
	AuthToken string `form:"authToken" json:"authToken" binding:"required"`
}

type APIGroupPushRequest struct {
	Msg       string `form:"msg" json:"msg" binding:"required"`
	ToGroupID string `form:"toGroupID" json:"toGroupID" binding:"required"`
	AuthToken string `form:"authToken" json:"authToken" binding:"required"`
}

type APIGroupCountRequest struct {
	GroupID string `form:"toGroupID" json:"toGroupID" binding:"required"`
}

type APIGroupInfoRequest struct {
	GroupID string `form:"toGroupID" json:"toGroupID" binding:"required"`
}

type APIResponse struct {
	Msg  string     `form:"msg" json:"msg"`
	Data common.Any `form:"data" json:"data"`
}

type APISignupRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type APISigninRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type APISignoutRequest struct {
	AuthToken string `form:"authToken" json:"authToken" binding:"required"`
}

type APIAuthCheckRequest struct {
	AuthToken string `form:"authToken" json:"authToken" binding:"required"`
}

type APISessionCheckRequest struct {
	AuthToken string `form:"authToken" json:"authToken" binding:"required"`
}
