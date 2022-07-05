package middlewares
// 处理App运行时的一些必要事件

import (
	"ginvel/common/helper"
	"ginvel/internal/reader/framework_toml"
	"ginvel/internal/processes/hold_request"
	"github.com/gin-gonic/gin"
	"time"
)

// Stat 捕捉接口运行耗时
// 此处使用的Next()请参考文档：https://blog.csdn.net/qq_37767455/article/details/104712028
func Stat(ctx *gin.Context) {
	start := float64(time.Now().UnixNano()) / 1000000 // ms
	ctx.Set("stat_start", helper.GetTimeDate("i.s.ms.ns")) // 设置公共参数。此参数范围是每个请求的公共参数，不是超全局参数，超全局参数请用globalData。

	// 等其他中间件先执行
	ctx.Next()

	// 获取运行耗时，ms
	end := float64(time.Now().UnixNano()) / 1000000
	latency := end - start

	// 设置公共参数（设置ctx每次请求的全局值）
	ctx.Set("stat_latency", latency)
	//fmt.Println("本次运行耗时=", latency, "ms")

	// 进入耗时治理服务
	holdRequest := hold_request.HoldRequest{}
	holdRequest.StatRequest(ctx)

	// 计时完成，中断所有后续函数调用
	ctx.Abort()
}

// HttpData 设置全局参数
func HttpData(ctx *gin.Context) {
	serverConfig := framework_toml.GetHttpServerConfig()
	host := serverConfig["Host"] + ":" + serverConfig["Port"]

	ctx.Set("host", host)

	ctx.Next()
}