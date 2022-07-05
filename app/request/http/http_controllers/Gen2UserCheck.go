package http_controllers

import (
	"ginvel/app/request/http/http_controllers/Gen2User"
	"ginvel/common/helper"
	"ginvel/common/kits"
	"github.com/gin-gonic/gin"
)

// Gen2UserCheck 接口安全校验
func Gen2UserCheck(ctx *gin.Context) {
	var state int64
	var msg string
	var err error
	var content interface{}

	appClass := kits.Input(ctx, "app_class")
	userID := kits.Input(ctx, "user_id")
	userName := kits.Input(ctx, "user_name")
	userToken := kits.Input(ctx, "user_token")

	classState := Gen2User.CheckAppClass(appClass)
	tokenState := Gen2User.CheckToken(userID, userName, userToken, appClass)

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
		msg = "非法参数，check，-2"
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
