package un_auth

import (
	"github.com/grin-ch/grin-api/api/grpc/captcha"
	"github.com/grin-ch/grin-auth/pkg/grin_error"
	"github.com/grin-ch/grin-auth/pkg/run/grpc"
	"github.com/grin-ch/grin-auth/pkg/run/http/ctx"
)

type SendCaptcha struct {
	ctx.PostCtx

	Contact string `binding:"required"`
	Purpose int    `binding:"required"`
}

func (act *SendCaptcha) Action() any {
	rsp, err := grpc.CaptchaClient.AsyncCode(act, &captcha.AsyncCodeReq{
		Contact: act.Contact,
		Purpose: captcha.Purpose(act.Purpose),
	})
	grin_error.GrpcOk(err)
	return rsp
}

func (act *SendCaptcha) Path() string {
	return "send_captcha"
}
