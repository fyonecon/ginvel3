package hold_framework

import (
	"fmt"
	"ginvel/app/arm/custom_interval"
	"ginvel/internal/processes/hold_hardware"
	"log"
)

// TimeInterval 全局定时器，默认精度30s/次
// 定时器ID，运行次数，周期
func (holdFramework *HoldFramework) TimeInterval(intervalId int, num int, timeout string) {
	var maxLog int = 1
	if num < maxLog { // 不必全部打印，只打印前几个即可
		if intervalId == 0 && num == 0 {
			log.Println("全局定时器初始化完成 >>>", " 定时器ID=", intervalId, " 运行次数num=", num, " 时间区间=", timeout)
		} else {
			log.Println("全局定时器已开启 >>>", " 定时器ID=", intervalId, " 运行次数num=", num, " 时间区间=", timeout)
		}
	} else if num == maxLog {
		log.Println("全局定时器日志显示已关闭，定时任务会继续运行。maxLog=", maxLog)
	} else {
		fmt.Print("\n\n") // 分割终端显示效果
		log.Println("全局定时器活动日志：num=", num)
	}

	// 其他定时任务
	// 硬件信息检测
	holdHardware := hold_hardware.HoldHardware{}
	holdHardware.Hardware(intervalId, num, timeout)
	// 热更服务
	holdFramework.HotUpdateRuler(num)
	// 其他自定义的全局定时任务
	custom_interval.CustomInterval(intervalId, num, timeout)

	fmt.Print("\n\n") // 分割终端显示效果
}
