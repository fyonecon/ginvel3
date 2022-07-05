package bootstrap

import (
	"ginvel/internal/middlewares"
	"ginvel/internal/processes/hold_framework"
	"ginvel/internal/reader/framework_toml"
	"ginvel/routes"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

// RunHttpServer 启动http服务，Powered by Gin。
func (bootstrap *Bootstrap) RunHttpServer() {
	var host string = bootstrap.FrameworkConfig.HttpServer.Host
	var port string = bootstrap.FrameworkConfig.HttpServer.Port
	var address string = host + ":" + port
	var routes = routes.Routes{}

	// 变量初始化
	httpConfig := framework_toml.GetHttpServerConfig()

	// Gin服务
	httpServer := gin.Default()

	// 捕捉接口运行参数（必须排第一）
	httpServer.Use(middlewares.Stat)

	// 设置全局ctx参数（必须排第二）
	httpServer.Use(middlewares.HttpData)

	// 拦截应用500报错，使之可视化
	httpServer.Use(middlewares.HttpError500)

	// 其他中间件
	// 拦截非法访问
	httpServer.Use(middlewares.HttpBlock)

	// Gin运行时：release、debug、test
	gin.SetMode(httpConfig["Env"])

	// 注册tpl路由
	routes.RouteHTML(httpServer)

	// 注册Api路由
	routes.RouteApi(httpServer)

	// 初始化定时器（立即运行定时器）
	holdFramework := hold_framework.HoldFramework{}
	holdFramework.TimeInterval(0, 0, "0")

	// 终端提示
	log.Println("Ginvel必要服务已经全部启动 >>> \n\n ")
	log.Println("http-api、ws接口访问地址示例：http://" + address)
	log.Println("http-html模板访问地址示例：http://" + address + "/html/example")
	log.Println("swagger访问地址示例：http://" + address + "/swagger/index.html \n\n ")

	err := httpServer.Run(address)
	if err != nil {
		log.Println("http服务遇到错误，运行中断，error：", err.Error())
		log.Println("提示：注意端口被占时应该首先更改对外暴露的端口，而不是微服务的端口。")
		os.Exit(200)
	}

	return
}