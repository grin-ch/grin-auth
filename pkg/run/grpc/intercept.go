package grpc

import (
	"github.com/grin-ch/grin-utils/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func recoveryFunc(p interface{}) (err error) {
	log.Logger.Errorf("%v", p)
	return status.Errorf(codes.Unknown, "%v", p)
}
