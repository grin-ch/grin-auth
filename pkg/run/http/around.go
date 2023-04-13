package http

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/grin-ch/grin-auth/cfg"
	"github.com/grin-ch/grin-auth/pkg/auth"
	"github.com/grin-ch/grin-auth/pkg/grin_error"
	"github.com/grin-ch/grin-auth/pkg/run/http/ctx"
	"github.com/grin-ch/grin-utils/log"
	"github.com/grin-ch/grin-utils/tool"
)

const (
	RequestID = "Requestid"
)

func Around(method string, act ctx.IAction) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		reqID := gctx.GetHeader(RequestID)
		if reqID == "" {
			reqID = strconv.FormatInt(tool.NewSnowFlakeID(), 10)
		}
		gctx.Header(RequestID, reqID)
		deadline := time.Duration(apiTimeOut(act.Module(), act.Path()))
		context, cancel := context.WithTimeout(gctx, deadline*time.Second)
		defer cancel()

		baseCtx := ctx.NewBaseCtx(context, gctx, getCliams(gctx))
		baseCtx.Set(ctx.Method, act.Method())
		baseCtx.Set(ctx.Module, act.Module())
		baseCtx.Set(ctx.Path, act.Path())
		defer func() {
			err := recover()
			if err != nil {
				act.ErrorHandle(err)
				gctx.Header("Content-Type", ctx.JSON)
				e, ok := err.(grin_error.IErr)
				if !ok {
					e = grin_error.UndefinedError(err)
				}
				gctx.JSON(200, deverErr(e))
			}
		}()
		act.Before(baseCtx)
		grin_error.ErrPanic(gctx.ShouldBind(act), grin_error.MissingParameter)
		func() {
			cost := tool.Cost()
			gctx.Header("Content-Type", act.ContextType())
			defer func() {
				rsp := act.Action()
				debugLog(act, reqID, rsp)
				switch act.ContextType() {
				case ctx.STRING:
					gctx.String(200, fmt.Sprintf("%v", rsp))
				case ctx.JSON:
					gctx.JSON(200, gin.H{
						"data": rsp,
						"cost": cost(),
					})
				default: //自由处理
				}
			}()
		}()
		act.After(baseCtx)
	}
}

func getCliams(gctx *gin.Context) *auth.Cliams {
	var c *auth.Cliams
	cliam, has := gctx.Get(auth.CliamsKey)
	if has {
		c = cliam.(*auth.Cliams)
	}
	return c
}

func apiTimeOut(module, path string) int {
	key := path
	if module != "" {
		key = module + "-" + key
	}
	if t, has := cfg.Config.Server.HttpServer.TimeoutAppoint[key]; has {
		return t
	}
	return cfg.Config.Server.HttpServer.Timeout
}

func deverErr(e grin_error.IErr) gin.H {
	rsp := gin.H{
		"code": e.Code(),
		"msg":  e.Msg(),
	}
	if cfg.Config.Server.HttpServer.Debug {
		rsp["err"] = e.Err()
	}
	return rsp
}

func debugLog(act ctx.IAction, reqID string, rsp any) {
	if cfg.Config.Server.HttpServer.Debug {
		log.Logger.Debugf("%s :%s IP:%s data:%+v", RequestID, reqID, act.ClientIP(), rsp)
	}
}
