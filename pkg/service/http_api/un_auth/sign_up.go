package un_auth

import (
	"github.com/grin-ch/grin-api/api/grpc/account"
	"github.com/grin-ch/grin-auth/pkg/grin_error"
	"github.com/grin-ch/grin-auth/pkg/run/grpc"
	"github.com/grin-ch/grin-auth/pkg/run/http/ctx"
)

type SignUp struct {
	ctx.PostCtx

	Nickname string `binding:"required"`
	Account  string `binding:"required"`
	Password string `binding:"required"`
	Captcha  string `binding:"required"`
}

func (act *SignUp) Action() any {
	rsp, err := grpc.UserClient.SignUp(act, &account.SignUpReq{
		Username: act.Nickname,
		Contact:  act.Account,
		Password: act.Password,
		Captcha:  act.Captcha,
	})
	grin_error.GrpcOk(err)
	return rsp
}

func (act *SignUp) Path() string {
	return "sign_up"
}
