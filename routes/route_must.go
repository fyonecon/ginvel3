package routes

// 默认路由，必要路由
// 预定义路由前缀
// web-api
// web-web
// websocket-ws

/*
  路由访问原则：统一用Any是遵循"宽进严出"的规则，在拦截器里面拦截（VerifyXXX.go、XXXCheck.go）具体请求事件。
  路由周期：请求路由名——>header过滤——>拦截请求频率——>校验请求方法和Token参数——>运行目标函数——>程序达到终点，关闭此次请求。
  路由写法 ：Any(路由名（必选）, header参数（可选）, 访问频率限制（可选）, 拦截器参数验证（可选）, 目标函数handler（必选）)。
  路由命名原则：推荐使用4层路由。第1层：api类还是web类；第2层：接口版本名；第3层：不同拦截器下的不同空间命名；第4层：目标函数handler。
*/

import (
	"ginvel/common/helper"
	"ginvel/common/kits"
	_ "ginvel/docs" // main.go根目录执行「swag init」生成的docs文件夹路径，_引包表示只执行init函数。
	"ginvel/internal/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// RouteMust ==系统必要路由，默認路由==
// handlerChain含义：
// 必选：middlewares.HttpCorsApi 输出json-header
// 可选：middlewares.HttpLimiterRate(10) 限流，数字代表每秒最多访问多少次就会禁止访问
// 可选：middlewares.HttpLimiterRanking(100) 熔断，数字代表CPU占用率(假设100%最大)到到xxx%就会熔断此接口
// 必选：Controllers.Welcome 目标函数，目标控制器到函数
func (routes *Routes) RouteMust(route *gin.Engine) {
	//var html string = frameworkConfig.View.ViewHTML // html文件目录
	var static string = frameworkConfig.View.ViewStatic // 静态文件主目录
	//var tpl string = frameworkConfig.View.ViewTPL // tpl文件主目录

	// 默认根路由
	route.Any("/", middlewares.HttpCorsApi, middlewares.HttpLimiterRate(2), middlewares.HttpLimiterRanking(100), func (ctx *gin.Context) {
		var url string = ctx.Request.Host + ctx.Request.URL.Path
		var ip string = ctx.ClientIP()
		kits.Log("默认路由接口 >>> " + url, ip)
		ctx.JSONP(http.StatusForbidden, gin.H{
			"state": 403,
			"msg": "Ginvel：默认路由接口",
			"content": map[string]interface{}{
				"time": helper.GetTimeDate("Ymd.His.ms.ns"),
				"tips": "此默认接口的耗时通常在40ms-100ms，大于500ms时应考虑参数赋值不完整。",
			},
		})
		return
	})

	// 404路由
	route.NoRoute(middlewares.HttpCorsApi, middlewares.HttpLimiterRate(4), func (ctx *gin.Context) {
		var url string = ctx.Request.Host + ctx.Request.URL.Path
		var ip string = ctx.ClientIP()
		kits.Error("404路由 >>> " + url, ip)
		ctx.JSONP(http.StatusNotFound, gin.H{
			"state": 404,
			"msg": "Ginvel：404，未定义此名称的路由名",
			"content": map[string]interface{}{
				"url": url,
				"time": helper.GetTimeDate("Ymd.His.ms.ns"),
			},
		})
		return
	})

	// 静态文件
	// ico图标
	route.StaticFile("/favicon.ico", static + "/img/favicon.ico")
	// robots文件，注意nginx规则里面不要拦截txt格式
	route.StaticFile("/robots.txt", static + "/robots.txt")

	// swagger接口文档，适配于Ginvel
	// 访问：127.0.0.1:9500/swagger/index.html
	route.GET("/swagger/*any", middlewares.HttpCorsWeb, middlewares.HttpLimiterRate(4), ginSwagger.WrapHandler(swaggerFiles.Handler))


	return
}
