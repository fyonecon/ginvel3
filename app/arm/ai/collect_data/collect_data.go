package collect_data
// 记录指标数据

import "fmt"

// CollectRequestStat 收集（记录）所有请求的耗时
// apiName：接口名或完整网址
// statLatency：0.00ms，接口请求耗时
// apiRanking：接口等级值
func CollectRequestStat(apiName string, statLatency interface{}, apiRanking interface{}) {
	if apiRanking != nil { // 排除无ranking的情况，如，静态文件接口是没有ranking的。
		fmt.Println("CollectRequestStat=", apiName, statLatency, apiRanking)
		// 保存数据
		// 待续

	}
	return
}

// CollectHardwareCPUPercent 收集（记录）所有服务器的CPU占有率
// serverIP：服务器的IP
// 定时器的计数
// cpuPercent：0.000000
func CollectHardwareCPUPercent(num int, timeout string, serverIP string, cpuPercent float64)  {
	fmt.Println("CollectHardwareCPUPercent=", num, timeout, serverIP, cpuPercent)
	// 保存数据
	// 待续

	return
}