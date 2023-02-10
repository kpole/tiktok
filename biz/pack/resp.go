package pack

import (
	"errors"
	"offer_tiktok/pkg/errno"
)

type BaseResp struct {
	StatusCode int32
	StatusMsg  string
}

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *BaseResp {
	return &BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
}
