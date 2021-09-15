package common

import "errors"

// todo 改为 const, error 分类 等等

var (
	InvaildUserNameError   = errors.New("invaild user name len == 0")
	InvaildPasswordError   = errors.New("invaild password len < 8")
	UnmatchedPasswordError = errors.New("unmatched password")
	UserHasExistedError    = errors.New("user has existed")
	SignoutFailedError     = errors.New("user signout failed")
)
