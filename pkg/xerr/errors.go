package xerr

import "github.com/zeromicro/x/errors"

// 错误工厂
func New(code int, msg string) error {
	return errors.New(code, msg)
}

// 自定义的错误信息输出
func NewMsg(msg string) error {
	return errors.New(SERVER_COMMON_ERROR, msg)
}

func NewMsgErr(msg string) error {
	return errors.New(SERVER_COMMON_ERROR, msg)
}

func NewCodeErr(code int) error {
	return errors.New(code, ErrMsg(code))
}

func NewInternalErr() error {
	return errors.New(SERVER_COMMON_ERROR, ErrMsg(SERVER_COMMON_ERROR))
}

func NewDBErr() error {
	return errors.New(DB_ERROR, ErrMsg(DB_ERROR))
}

func NewReqParamErr() error {
	return errors.New(REQUEST_PARAM_ERROR, ErrMsg(REQUEST_PARAM_ERROR))
}
