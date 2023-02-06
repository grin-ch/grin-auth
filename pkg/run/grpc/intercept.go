package grpc

import (
	"runtime"

	"github.com/grin-ch/grin-utils/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const defaultStackSize = 4096

func recoveryFunc(p interface{}) (err error) {
	var buf [defaultStackSize]byte
	n := runtime.Stack(buf[:], false)
	log.Logger.Errorf("err:%v trace: %s", p, buf[:n])
	return status.Errorf(codes.Unknown, "%v", p)
}
