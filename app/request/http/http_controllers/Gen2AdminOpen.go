package http_controllers

import (
	"ginvel/app/request/http/http_controllers/Gen2Admin"
	"ginvel/common/kits"
	"github.com/gin-gonic/gin"
)

// Gen2AdminOpen 接口安全校验
func Gen2AdminOpen(ctx *gin.Context) {
	var state int64
	var msg string
	var err error
	var content interface{}

	appClass := kits.Input(ctx, "app_class")
	classState := Gen2Admin.CheckAppClass(appClass)

	if !classState {
		state = 0
		msg = "非法参数，open，-1"
		content = ""
		ctx.JSONP(200, map[string]interface{}{
			"state": state,
			"msg": msg,
			"content": map[string]interface{}{
				"err": err,
				"content": content,
			},
		})
		ctx.Abort()
	}else {
		ctx.Next()
	}

	return
}