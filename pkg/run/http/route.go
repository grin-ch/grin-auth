package http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/grin-ch/grin-auth/pkg/auth"
	"github.com/grin-ch/grin-auth/pkg/grin_error"
	"github.com/grin-ch/grin-auth/pkg/run/http/ctx"
	"github.com/grin-ch/grin-auth/pkg/service/http_api/un_auth"
)

var router *gin.Engine

func Router(mode string) *gin.Engine {
	gin.SetMode(mode)
	router = gin.New()
	router.Use(cors())
	unAuthApi()
	router.Use(middlewares()...)

	return router
}

func cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		origin := ctx.Request.Header.Get("Origin")
		if origin != "" {
			ctx.Header("Access-Control-Allow-Origin", "*") // Access-Control-Allow-Origin * 替换为指定的域名
			ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			ctx.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			ctx.JSON(http.StatusOK, "ok!")
		}
		ctx.Next()
	}
}

func middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		authBase,
	}
}

func authBase(ctx *gin.Context) {
	token := ctx.GetHeader(auth.TokenKey)
	if token != "" {
		token = strings.Replace(token, "Bearer ", "", 1)
		cliams, err := auth.ParseJWT(token)
		if err == nil {
			ctx.Set(auth.CliamsKey, cliams)
			ctx.Next()
			return
		}
	}

	ctx.JSON(http.StatusOK, grin_error.EnumData(grin_error.AuthError))
	ctx.Abort()
}

func registery(act ctx.IAction) {
	var url string
	if act.Module() == "" {
		url = fmt.Sprintf("/%s", act.Path())
	} else {
		url = fmt.Sprintf("/%s/%s", act.Module(), act.Path())
	}

	router.Handle(act.Method(), url, Around(act.Method(), act))
}

// 无需验证身份的接口
func unAuthApi() {
	registery(&un_auth.Health{})
	registery(&un_auth.SendCaptcha{})
	registery(&un_auth.SignUp{})
	registery(&un_auth.SignInCode{})
	registery(&un_auth.SignIn{})
}
