package grin_error

type errCode int
type errMsg string

type IErr interface {
	Code() errCode
	Msg() errMsg
	Err() string
}

type errAble struct {
	code errCode
	msg  errMsg
	err  error
}

func (err *errAble) Code() errCode {
	return err.code
}

func (err *errAble) Msg() errMsg {
	return err.msg
}

func (err *errAble) Err() string {
	if err.err == nil {
		return ""
	}
	return err.err.Error()
}
