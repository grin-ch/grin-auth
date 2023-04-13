package un_auth

import (
	"github.com/grin-ch/grin-api/api/grpc/account"
	"github.com/grin-ch/grin-auth/pkg/grin_error"
	"github.com/grin-ch/grin-auth/pkg/run/grpc"
	"github.com/grin-ch/grin-auth/pkg/run/http/ctx"
)

type SignIn struct {
	ctx.PostCtx

	CaptchaKey string
	Captcha    string
}

func (act *SignIn) Action() any {
	contact, passwd, has := act.GinCtx().Request.BasicAuth()
	grin_error.PanicWhen(!has, grin_error.MissingParameter)

	rsp, err := grpc.UserClient.SignIn(act, &account.SignInReq{
		Contact:  contact,
		Password: passwd,
	})
	grin_error.GrpcOk(err)
	return rsp
}

func (act *SignIn) Path() string {
	return "sign_in"
}
