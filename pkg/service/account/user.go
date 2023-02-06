package account

import (
	"context"

	"github.com/grin-ch/grin-api/api/grpc/account"
	"github.com/grin-ch/grin-api/api/grpc/captcha"
	"github.com/grin-ch/grin-auth/pkg/auth"
	"github.com/grin-ch/grin-auth/pkg/model"
	"github.com/grin-ch/grin-auth/pkg/util"
	"github.com/grin-ch/grin-utils/log"
	"github.com/grin-ch/grin-utils/tool"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userService struct {
	*dbServer
	expires int
	signed  string
	issuer  string

	captchaClient captcha.CaptchaServiceClient
}

// NewUserService  userService impl
func NewUserService(dbClient *model.Client, expires int, signed, issuer string,
	captchaClient captcha.CaptchaServiceClient) account.UserServiceServer {
	return &userService{
		dbServer:      newDbServer(dbClient),
		expires:       expires,
		signed:        signed,
		issuer:        issuer,
		captchaClient: captchaClient,
	}
}

// AuthFuncOverride 覆盖authFunc
// 跳过部分接口的权限验证
func (*userService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

// 注册
func (s *userService) SignUp(ctx context.Context, req *account.SignUpReq) (*account.SignUpRsp, error) {
	// 参数校验
	if len(req.Username) < 3 || len(req.Username) > 24 {
		return nil, status.Errorf(codes.InvalidArgument, "username must be between 3 and 12 character")
	}
	var phone, email string
	if util.ValidatePhoneNumber(req.Contact) {
		phone = req.Contact
	} else if util.ValidateEmail(req.Contact) {
		email = req.Contact
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "contact invalid")
	}
	if len(req.Password) < 8 {
		return nil, status.Errorf(codes.InvalidArgument, "password length less 8")
	}
	encodePwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		log.Logger.Errorf("password encode err:%v", err)
		return nil, status.Errorf(codes.InvalidArgument, "params invalid")
	}

	// 验证码校验
	rsp, err := s.captchaClient.Verify(ctx, &captcha.VerifyReq{
		Key:     req.Contact,
		Value:   req.Captcha,
		Purpose: captcha.Purpose_SIGN_UP,
	})
	if err != nil {
		log.Logger.Errorf("captcha verify err:%v", err)
		return nil, status.Errorf(codes.Internal, "captcha verify error")
	}
	if !rsp.Success {
		return nil, status.Errorf(codes.InvalidArgument, "captcha invalid")
	}

	// 创建帐号
	_, err = s.CreateUser(ctx, req.Username, phone, email, string(encodePwd))
	if err != nil {
		log.Logger.Errorf("CreateUser err:%v", err)
		return nil, status.Errorf(codes.AlreadyExists, "username or contact already exists")
	}
	return &account.SignUpRsp{
		Success: true,
		Message: "sign up!",
	}, nil
}

// 登入
func (s *userService) SignIn(ctx context.Context, req *account.SignInReq) (*account.SignInRsp, error) {
	// 验证码校验
	rsp, err := s.captchaClient.Verify(ctx, &captcha.VerifyReq{
		Key:     req.CaptchaKey,
		Value:   req.Captcha,
		Purpose: captcha.Purpose_SIGN_IN,
	})
	if err != nil {
		log.Logger.Errorf("captcha verify err:%v", err)
		return nil, status.Errorf(codes.Internal, "captcha verify error")
	}
	if !rsp.Success {
		return nil, status.Errorf(codes.InvalidArgument, "captcha invalid")
	}

	// 密码校验
	if len(req.Contact) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "contact invalid")
	}
	user, err := s.FindUserByAccount(ctx, req.Contact)
	if err != nil {
		log.Logger.Errorf("FindUserByAccount err:%v", err)
		return nil, status.Errorf(codes.Unauthenticated, "account or password invalid")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "account or password invalid")
	}

	token, err := auth.GenerateJWT(s.expires, s.signed, s.issuer, auth.RoleBase{
		Id:       user.ID,
		UUID:     tool.MustUUIDv4(),
		Avatar:   user.Edges.UserData.AvatarURL,
		Nickname: user.Edges.UserData.Nickname,
		Sex:      user.Edges.UserData.Sex.String(),
	})
	if err != nil {
		log.Logger.Errorf("generate token err:%v", err)
		return nil, status.Errorf(codes.Internal, "generate token error")
	}

	return &account.SignInRsp{
		Token: token,
	}, nil
}

// 重置密码
func (s *userService) ResetPasswd(ctx context.Context, req *account.ResetPasswdReq) (*account.ResetPasswdRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "Unimplemented")
}
