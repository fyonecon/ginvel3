package Gen2Admin

import (
	"ginvel/common/kits"
	"github.com/gin-gonic/gin"
)


func AdminTest1(ctx *gin.Context)  {
	var str string = "abc#@9(%$19#@xn34-333"
	var str1 string = kits.Encode(str,"abcL")
	var str2 string = kits.Decode(str1, "abcL")

	// 接口返回
	back := map[string]interface{}{
		"state": 1,
		"msg": "test",
		"content": []interface{}{
			str,
			str1,
			str2,
		},
	}
	ctx.JSONP(200, back)
	return
}
