package grpc

import (
	"context"
	"runtime"

	"github.com/grin-ch/grin-auth/pkg/auth"
	"github.com/grin-ch/grin-utils/log"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	bearer           = "bearer"
	defaultStackSize = 4096
)

func recoveryFunc(p interface{}) (err error) {
	var buf [defaultStackSize]byte
	n := runtime.Stack(buf[:], false)
	log.Logger.Errorf("err:%v trace: %s", p, buf[:n])
	return status.Errorf(codes.Unknown, "%v", p)
}

func authFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, bearer)
	if err != nil {
		return nil, err
	}
	claims, err := auth.ParseJWT(token)
	if err != nil {
		return nil, err
	}
	err = claims.Valid()
	if err != nil {
		return nil, err
	}
	return ctx, nil
}
