package main_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/grin-ch/grin-api/api/grpc/account"
	"github.com/grin-ch/grin-auth/cfg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	client account.UserServiceClient
)

func TestMain(m *testing.M) {
	cfg.InitConfig()

	host := fmt.Sprintf("%s:%d", cfg.Config.Server.Host, cfg.Config.Server.AccountServer.Info.Port)
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//创建endpoint，并指明grpc调用的接口名和方法名
	client = account.NewUserServiceClient(conn)
	m.Run()
}

func TestSignUp(t *testing.T) {
	rsp, err := client.SignUp(context.Background(), &account.SignUpReq{
		Username: "nickname",
		Contact:  "email@email.com",
		Password: "password",
		Captcha:  "captcha code",
	})

	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", rsp)
}

func TestSignIn(t *testing.T) {
	rsp, err := client.SignIn(context.Background(), &account.SignInReq{
		Contact:    "email@email.com",
		Password:   "password",
		CaptchaKey: "fa3e39acbdd38d3917e19cd2a4e7f54e",
		Captcha:    "FFHF",
	})

	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", rsp)
}
