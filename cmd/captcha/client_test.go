package main_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/grin-ch/grin-api/api/grpc/captcha"
	"github.com/grin-ch/grin-auth/cfg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	client captcha.CaptchaServiceClient
)

func TestMain(m *testing.M) {
	cfg.InitConfig()

	host := fmt.Sprintf("%s:%d", cfg.Config.Server.Host, cfg.Config.Server.CaptchaServer.Info.GrpcPort)
	fmt.Println(host)
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//创建endpoint，并指明grpc调用的接口名和方法名
	client = captcha.NewCaptchaServiceClient(conn)
	m.Run()
}

func TestAsyncCode(t *testing.T) {
	rsp, err := client.AsyncCode(context.Background(), &captcha.AsyncCodeReq{
		Contact: "email@email.com",
		Purpose: captcha.Purpose_SIGN_UP,
	})

	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", rsp)
}

func TestGraphCaptcha(t *testing.T) {
	rsp, err := client.GraphCaptcha(context.Background(), &captcha.GraphCaptchaReq{
		Purpose: captcha.Purpose_SIGN_IN,
	})

	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", rsp)
}
