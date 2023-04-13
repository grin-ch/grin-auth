package un_auth

import (
	"github.com/grin-ch/grin-api/api/grpc/captcha"
	"github.com/grin-ch/grin-auth/pkg/grin_error"
	"github.com/grin-ch/grin-auth/pkg/run/grpc"
	"github.com/grin-ch/grin-auth/pkg/run/http/ctx"
)

type SignInCode struct {
	ctx.GetCtx
}

func (act *SignInCode) Action() any {
	rsp, err := grpc.CaptchaClient.GraphCaptcha(act, &captcha.GraphCaptchaReq{
		Purpose: captcha.Purpose_SIGN_IN,
	})
	grin_error.GrpcOk(err)
	return rsp
}

func (act *SignInCode) Path() string {
	return "sign_in_code"
}
