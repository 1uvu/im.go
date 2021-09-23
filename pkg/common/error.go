package common

import (
	"errors"
)

// todo error 分类, 实现包装器 等等
type CommonError error

var (
	ErrInvaildUserName = errors.New("invaild user name len == 0")
	ErrInvaildPassword = errors.New("invaild password len < 8")

	ErrUnmatchedPassword  = errors.New("unmatched password")
	ErrUnmatchedAuthToken = errors.New("unmatched auth token")

	ErrSessionHasExpired  = errors.New("session has expired")
	ErrSessionDeletFailed = errors.New("session delete failed")

	ErrUserServerQuitFailed = errors.New("user quit server failed")
	ErrUserHasExisted       = errors.New("user has existed")
	ErrUserNotExisted       = errors.New("user not existed")
	ErrUserSignoutFailed    = errors.New("user signout failed")

	ErrMarshalPushArgFailed = errors.New("marshal push arg failed")
	ErrPublishFailed        = errors.New("publish arg failed")

	ErrGetGroupUsersFailed = errors.New("get group user infos failed")
	ErrGetGroupCountFailed = errors.New("get group user count failed")
	ErrGroupIsNotLive      = errors.New("group has 0 user and it is not live now")

	ErrConnectFailed    = errors.New("connect failed")
	ErrDisconnectFailed = errors.New("disconnect failed")

	ErrNaNRPCArg = errors.New("rpc arg is nil")
)

// func Err() CommonError {
// 	return
// }

// func ErrWithMsg(err CommonError, errMsg string) CommonError {
// 	return fmt.Errorf("error: %s with msg: %s", err.Error(), errMsg)
// }
