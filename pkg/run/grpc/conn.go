package grpc

import (
	"github.com/grin-ch/grin-api/api/grpc/account"
	"github.com/grin-ch/grin-api/api/grpc/captcha"
	"github.com/grin-ch/grin-auth/cfg"
	"google.golang.org/grpc"
)

var (
	UserClient    account.UserServiceClient
	CaptchaClient captcha.CaptchaServiceClient
)

// 初始化客户端
func initClients(resolve func(string) (*grpc.ClientConn, error), connectors ...ClientConnector) {
	for _, c := range connectors {
		name, fn := c()
		initClient(resolve, name, fn)
	}
}

func initClient(resolve func(string) (*grpc.ClientConn, error),
	serverName string, fn func(*grpc.ClientConn)) {
	conn, err := resolve(serverName)
	if err != nil {
		panic(err)
	}
	fn(conn)
}

type ClientConnector func() (string, func(conn *grpc.ClientConn))

func AllConn() []ClientConnector {
	return []ClientConnector{
		CaptchaConn,
		UserConn,
	}
}

func CaptchaConn() (string, func(*grpc.ClientConn)) {
	name := cfg.Config.Server.CaptchaServer.Info.Name
	return name, func(cc *grpc.ClientConn) {
		CaptchaClient = captcha.NewCaptchaServiceClient(cc)
	}
}

func UserConn() (string, func(*grpc.ClientConn)) {
	name := cfg.Config.Server.AccountServer.Info.Name
	return name, func(cc *grpc.ClientConn) {
		UserClient = account.NewUserServiceClient(cc)
	}
}
