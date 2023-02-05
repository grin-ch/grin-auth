package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func recoveryFunc(p interface{}) (err error) {
	return status.Errorf(codes.Unknown, "%v", p)
}
