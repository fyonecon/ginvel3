package routes_ws

import (
	"ginvel/app/request/ws/ws_controllers"
	"ginvel/common/helper"
	"ginvel/internal/middlewares"
	"ginvel/internal/reader/framework_toml"
	"github.com/gin-gonic/gin"
)

// DefaultRoute 分组路由
func DefaultRoute(route *gin.Engine)  {
	var htmlPath string = helper.FrameInfo["framework_path"] + framework_toml.GetViewConfig()["ViewHTML"] // html文件目录
	//var staticPath string = helper.FrameInfo["framework_path"] + framework_toml.GetViewConfig()["ViewStatic"] // 静态文件目录

	// 示例-html读取websocket接口示例
	route.StaticFile("/ws", htmlPath + "/pages/example-ws.html")

	// api分组路由
	ws := route.Group("/ws", middlewares.HttpCorsApi)
	{
		ws.Any("/ping1", ws_controllers.Ping1)
	}



}
