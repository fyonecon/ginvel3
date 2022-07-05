package routes
// 按分组注册用户路由，基于Gin，仅适用于http和websocket。

import (
	"fmt"
	"ginvel/common/helper"
	"ginvel/internal/reader/framework_toml"
	"ginvel/routes/routes_api"
	"ginvel/routes/routes_api/Gen2Routes"
	"ginvel/routes/routes_html"
	"ginvel/routes/routes_ws"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

type Routes struct {}

var frameworkConfig = framework_toml.ReadFrameworkToml()

// RouteHTML 注册html模板相关、分组模板
func (routes *Routes) RouteHTML(route *gin.Engine)  {
	//var html string = frameworkConfig.View.ViewHTML // html文件目录
	var static string = frameworkConfig.View.ViewStatic // 静态文件主目录
	var tpl string = frameworkConfig.View.ViewTPL // tpl文件主目录

	if len(static) == 0 || len(tpl) == 0 {
		helper.Log("static或tpl的地址参数不能为空，运行中断。")
		fmt.Println("static或tpl的地址参数不能为空，运行中断。", "static="+static, "tpl="+tpl)
		os.Exit(200)
	}
	log.Println("运行自定义注册html路由文件 >>> ")

	// 必要多模板基础参数
	// 静态文件目录，文档：https://www.tizi365.com/archives/271.html
	route.Static("/static", static)
	// 注册html模板，文档：https://www.tizi365.com/archives/268.html
	route.LoadHTMLGlob(tpl)
	// 初始默认路由
	routes_html.DefaultRoute(route)

	// 注册其他html路由
	//


	return
}

// RouteApi 注册api接口相关、分组接口
func (routes *Routes) RouteApi(route *gin.Engine)  {
	log.Println("运行自定义注册api、ws路由文件 >>> ")

	// 必要路由：处理默认路由、静态文件路由、404路由等
	routes.RouteMust(route)
	// 初始默认路由
	routes_api.DefaultRoute(route)
	routes_ws.DefaultRoute(route)

	// 注册其他api、ws路由
	//
	Gen2Routes.Gen2AdminApi(route)


	return
}