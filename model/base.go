package model

import "fmt"

const (
	successCode = 0
)

type Errorer interface {
	Error() error
}

func NewCodeMsg(code int32, msg string) CodeMsg {
	cm := CodeMsg{
		Code: code,
		Msg:  msg,
	}
	cm.Err = cm.Error()
	return cm
}

func NewCodeMsgWithErr(err error) CodeMsg {
	return CodeMsg{Err: err}
}

type CodeMsg struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Err  error  `json:"err,omitempty"`
}

func (c *CodeMsg) Error() error {
	if c.Code == successCode {
		return nil
	}
	return fmt.Errorf("code: %d, msg: %s", c.Code, c.Msg)
}

type WebResp struct {
	CodeMsg
	Data interface{} `json:"data,omitempty"`
}
