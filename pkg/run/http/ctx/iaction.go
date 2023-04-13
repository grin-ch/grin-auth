package ctx

import (
	"net/http"
	"runtime"

	"github.com/grin-ch/grin-utils/log"
)

const (
	defaultStackSize = 4096

	Method = "Method"
	Module = "Module"
	Path   = "Path"
)

type IAction interface {
	ICtx
	Before(ICtx)
	Action() any
	After(ICtx)
	ErrorHandle(any)
	Method() string
	Module() string
	Path() string
	ContextType() string
}

type GetCtx struct{ baseAction }

func (GetCtx) Method() string { return http.MethodGet }

type PutCtx struct{ baseAction }

func (PutCtx) Method() string { return http.MethodPut }

type PostCtx struct{ baseAction }

func (PostCtx) Method() string { return http.MethodPost }

type DelCtx struct{ baseAction }

func (DelCtx) Method() string { return http.MethodDelete }

type baseAction struct {
	ICtx

	method string
	module string
	path   string
}

func (ctx *baseAction) Before(ictx ICtx) {
	ctx.ICtx = ictx
	ctx.method = getStr(ictx, Method)
	ctx.module = getStr(ictx, Module)
	ctx.path = getStr(ictx, Path)
}
func (*baseAction) Action() any { return nil }
func (*baseAction) After(ICtx)  {}
func (ctx *baseAction) ErrorHandle(err any) {
	var buf [defaultStackSize]byte
	n := runtime.Stack(buf[:], false)
	log.Logger.Errorf("Module:%s,Path:%s, Error: %v\nTrace: %s", ctx.Module(), ctx.Path(), err, buf[:n])
}
func (ctx *baseAction) Method() string      { return ctx.method }
func (ctx *baseAction) Module() string      { return ctx.module }
func (ctx *baseAction) Path() string        { return ctx.path }
func (ctx *baseAction) ContextType() string { return JSON }

func getStr(ictx ICtx, key string) string {
	val, has := ictx.Get(key)
	if !has {
		return ""
	}
	return val.(string)
}
