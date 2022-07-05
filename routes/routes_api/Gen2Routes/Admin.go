package Gen2Routes

import (
	"ginvel/app/request/http/http_controllers"
	"ginvel/app/request/http/http_controllers/Gen2Admin"
	"ginvel/internal/middlewares"
	"github.com/gin-gonic/gin"
)

// Gen2AdminApi 分组路由
func Gen2AdminApi(route *gin.Engine)  {
	// api分组路由
	api := route.Group("/api/", middlewares.HttpCorsApi)
	{
		gen2 := api.Group("/gen2/")
		{
			open := gen2.Group("/admin/", middlewares.HttpLimiterRate(2), middlewares.HttpLimiterRanking(90), http_controllers.Gen2AdminOpen)
			{
				open.Any("admin_login", Gen2Admin.AdminLogin)
				open.Any("level_admin", Gen2Admin.LevelAdmin)
				open.Any("make_captcha", Gen2Admin.MakeCaptcha)

				// test
				open.Any("admin_test1", Gen2Admin.AdminTest1)

			}

			check := gen2.Group("/admin/", middlewares.HttpLimiterRate(2), middlewares.HttpLimiterRanking(90), http_controllers.Gen2AdminCheck)
			{
				check.Any("check_login_token", Gen2Admin.CheckLoginToken)

				check.Any("list_admin", Gen2Admin.ListAdmin)
				check.Any("that_admin", Gen2Admin.ThatAdmin)
				check.Any("add_admin", Gen2Admin.AddAdmin)
				check.Any("update_admin", Gen2Admin.UpdateAdmin)
				check.Any("del_admin", Gen2Admin.DelAdmin)

			}

		}
	}
}

