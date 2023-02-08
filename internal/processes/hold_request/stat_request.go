package hold_request

// http耗时服务参数

import (
	"fmt"
	"ginvel/app/arm/collect_data"
	"ginvel/common/helper"
	"ginvel/common/kits"
	"github.com/gin-gonic/gin"
)

type HoldRequest struct{}

// StatRequest 计算每个请求的相关参数
func (holdRequest *HoldRequest) StatRequest(ctx *gin.Context) {

	// 当前请求链接
	_host, _ := ctx.Get("host")
	host := helper.ValueInterfaceToString(_host)
	uri := host + ctx.Request.RequestURI // 完整的请求链接
	// 获取ctx每次请求的全局值
	_statLatency, _ := ctx.Get("stat_latency") // 请求耗时
	_apiRanking, _ := ctx.Get("api_ranking")   // 请求api的ranking值

	fmt.Println("Stat@请求uri=", uri, "接口耗时=", _statLatency, "ms", "，接口等级=", _apiRanking)

	// 收集指标（将每个请求）
	collect_data.CollectRequestStat(uri, _statLatency, _apiRanking)

	// 特殊情况记录
	statLatency := helper.StringToFloat(helper.ValueInterfaceToString(_statLatency)) // ms
	if statLatency > 4*1000 {                                                        // 超过4s都记录下来
		kits.Error(helper.ValueInterfaceToString(_statLatency)+"ms；"+uri, ctx.ClientIP())
	}

	ctx.Next()
}
