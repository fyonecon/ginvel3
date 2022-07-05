package routes_html

import (
	"ginvel/common/helper"
	"ginvel/internal/middlewares"
	"ginvel/internal/reader/framework_toml"
	"github.com/gin-gonic/gin"
)

// DefaultRoute 分组路由
func DefaultRoute(route *gin.Engine)  {
	var htmlPath string = helper.FrameInfo["framework_path"] + framework_toml.GetViewConfig()["ViewHTML"] // html文件目录
	//var staticPath string = helper.FrameInfo["framework_path"] + framework_toml.GetViewConfig()["ViewStatic"] // 静态文件目录

	// 示例-直接输出HTML文件。http://127.0.0.1:9500/html
	route.StaticFile("/html", htmlPath + "/welcome.html")
	route.StaticFile("/html/welcome", htmlPath + "/welcome.html")
	// 示例-直接输出HTML文件。http://127.0.0.1:9500/html/vue3
	route.StaticFile("/html/vue3", htmlPath + "/vue3.html")

	// 示例-多html渲染模板
	html := route.Group("/html", middlewares.HttpCorsWeb, middlewares.HttpLimiterRate(30), middlewares.HttpLimiterRanking(80))
	{
		// 直接输出HTML文件。http://127.0.0.1:9500/html/welcome
		html.GET("/:file", func (ctx *gin.Context) {
			var file string = ctx.Param("file")
			html.StaticFile("/"+file, htmlPath + file+".html")
			return
		})

		// 模板输出HTML。http://127.0.0.1:9500/html/example
		html.GET("/example", func (ctx *gin.Context) {
			var url string = ctx.Request.Host + ctx.Request.URL.Path
			var ip string = ctx.ClientIP()

			var tpl string = "pages/example-tpl.tmpl"
			var data map[string]interface{} = map[string]interface{}{
				"url": url,
				"ip": ip,
			}
			ctx.HTML(200, tpl, data)
			return
		})

	}

}
