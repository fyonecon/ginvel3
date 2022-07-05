package http_controllers

import (
	"ginvel/app/request/http/http_controllers/Gen2Admin"
	"ginvel/common/helper"
	"ginvel/common/kits"
	"github.com/gin-gonic/gin"
)

// Gen2AdminCheck 接口安全校验
func Gen2AdminCheck(ctx *gin.Context) {
	var state int64
	var msg string
	var err error
	var content interface{}

	appClass := kits.Input(ctx, "app_class")
	loginID := kits.Input(ctx, "login_id")
	loginName := kits.Input(ctx, "login_name")
	loginToken := kits.Input(ctx, "login_token"); loginToken = helper.DecodeBase64(loginToken)

	classState := Gen2Admin.CheckAppClass(appClass)
	tokenState := Gen2Admin.CheckToken(loginID, loginName, loginToken, appClass)

	if classState {
		state = helper.InterfaceToInt(tokenState["state"])
		msg = helper.InterfaceToString(tokenState["msg"])
		content = tokenState["content"]
		if state != 1 {
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
	}else {
		state = 0
		msg = "非法参数，check，-1"
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
	}

	return
}

