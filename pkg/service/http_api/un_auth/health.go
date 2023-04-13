package un_auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grin-ch/grin-auth/pkg/run/http/ctx"
)

type Health struct {
	ctx.GetCtx
}

func (act *Health) Action() any {
	now := time.Now()
	return gin.H{
		"time":    now.Format("2006-01-02 15:04:05"),
		"weekday": now.Weekday().String(),
	}
}

func (act *Health) Path() string {
	return "health"
}
