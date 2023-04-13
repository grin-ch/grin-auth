package grin_error

import (
	"fmt"
)

var (
	codeEnumMap = map[errCode]errMsg{}
)

func setErrEnum(code errCode, msg errMsg) {
	if _, has := codeEnumMap[code]; has {
		panic(fmt.Sprintf("duplicate error code:%d", code))
	}
	codeEnumMap[code] = msg
}

func UndefinedError(err any) *errAble {
	e := &errAble{
		code: -1,
		msg:  "UndefinedError",
	}
	if err != nil {
		e.err = fmt.Errorf("%v", err)
	}
	return e
}

func EnumData(code errCode, args ...interface{}) interface{} {
	msg, has := codeEnumMap[code]
	if !has {
		msg = UndefinedError(nil).msg
	}
	if len(args) > 0 {
		msg = errMsg(fmt.Sprintf(string(msg), args...))
	}
	return map[string]any{
		"code": code,
		"msg":  msg,
	}
}

func ErrPanic(err error, code errCode, args ...interface{}) {
	if err == nil {
		return
	}
	var e *errAble
	msg, has := codeEnumMap[code]
	if !has {
		e = UndefinedError(err)
	} else {
		e = &errAble{
			code: code,
			msg:  msg,
			err:  err,
		}
	}
	if len(args) > 0 {
		e.msg = errMsg(fmt.Sprintf(string(e.msg), args...))
	}
	panic(e)
}

func PanicWhen(doPanic bool, code errCode, args ...interface{}) {
	if doPanic {
		handlePanic(code, args...)
	}
}

func handlePanic(code errCode, args ...interface{}) {
	var e *errAble
	msg, has := codeEnumMap[code]
	if !has {
		e = UndefinedError(nil)
	} else {
		e = &errAble{
			code: code,
			msg:  msg,
		}
	}
	if len(args) > 0 {
		e.msg = errMsg(fmt.Sprintf(string(e.msg), args...))
	}
	panic(e)
}

func GrpcOk(err error) {
	if err != nil {
		panic(err)
	}
}
