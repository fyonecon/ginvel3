package routes_api

import (
	"ginvel/common/helper"
	"ginvel/common/kits"
	"ginvel/internal/middlewares"
	"github.com/gin-gonic/gin"
)

// DefaultRoute 分组路由
func DefaultRoute(route *gin.Engine)  {
	//var htmlPath string = helper.FrameInfo["framework_path"] + framework_toml.GetViewConfig()["ViewHTML"] // html文件目录
	//var staticPath string = helper.FrameInfo["framework_path"] + framework_toml.GetViewConfig()["ViewStatic"] // 静态文件目录

	// api_json数据输出
	route.Any("/api", middlewares.HttpCorsApi, middlewares.HttpLimiterRate(50), middlewares.HttpLimiterRanking(100), func (ctx *gin.Context) {
		var url string = ctx.Request.Host + ctx.Request.URL.Path
		var ip string = ctx.ClientIP()
		kits.Log("api根路由 >>> " + url, ip)
		ctx.JSONP(200, gin.H{
			"state": 1,
			"msg": "api根路由",
			"content": map[string]interface{}{
				"url": url,
				"time": helper.GetTimeDate("Ymd.His.ms.ns"),
			},
		})
		return
	})

}

