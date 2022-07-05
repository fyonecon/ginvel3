package middlewares

import (
	"ginvel/common/helper"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
	"time"
)

// HttpLimiterRate 拦截http请求频率，请求限流，num/s
func HttpLimiterRate(_max float64) gin.HandlerFunc {
	var max float64
	var tbOptions limiter.ExpirableOptions
	tbOptions.DefaultExpirationTTL = time.Second // 默认按每秒

	if _max > 0 || _max <= 1000 {
		max = _max
	}else {
		max = 20
	}
	lmt := tollbooth.NewLimiter(max, &tbOptions) // 默认4次/秒，建议范围[1，40]

	return func(ctx *gin.Context) {
		httpError := tollbooth.LimitByRequest(lmt, ctx.Writer, ctx.Request)
		if httpError != nil {
			//ctx.Data(httpError.StatusCode, lmt.GetMessageContentType(), []byte(httpError.Message))
			ctx.JSON(429, gin.H{
				"state": 429,
				"msg": "触及接口访问频率限制",
				"content": gin.H{
					"gv_version": helper.FrameInfo["gv_version"],
					"go_version": helper.FrameInfo["go_version"],
					//"ip": ctx.ClientIP(),
				},
			})

			ctx.Abort()
		} else {
			ctx.Next()
		}
	}
}
