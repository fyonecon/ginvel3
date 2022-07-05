package middlewares
// 接管服务器500错误，使错误可视化

import (
	"fmt"
	"ginvel/common/helper"
	"ginvel/common/kits"
	"github.com/gin-gonic/gin"
	"runtime"
	"strings"
)

// HideHttpInfo 隐藏必要关键词
func HideHttpInfo(info string) string {

	// Ginvel
	var hideElement = []string{
		"go",
		"/",
		".",
		"ginlaravel",
		"ginvel",
		"app",
		"framework",
		"bootstrap",
		"helper",
		"runtime",
		"request",
		"kits",
		"toml",
		"frame",
		"providers",
	}
	for _, value := range hideElement{
		info = strings.Replace(info, value, "*", -1)
	}

	return info
}

// HttpError500 抛出500错误
func HttpError500 (ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			//打印错误堆栈信息
			fmt.Printf("\n Http panic: %v \n\n", err)
			// debug.PrintStack() // 显示报错详情
			kits.Error(helper.ValueInterfaceToString(err), ctx.ClientIP())

			pc := make([]uintptr, 8)
			runtime.Callers(2, pc)
			//f := runtime.FuncForPC(pc[0])

			_fc0 := runtime.FuncForPC(pc[0]).Name()
			_fc1 := runtime.FuncForPC(pc[1]).Name()
			_fc2 := runtime.FuncForPC(pc[2]).Name()
			_fc3 := runtime.FuncForPC(pc[3]).Name()
			_fc4 := runtime.FuncForPC(pc[4]).Name()
			_fc5 := runtime.FuncForPC(pc[5]).Name()
			_fc6 := runtime.FuncForPC(pc[6]).Name()

			fc0 := HideHttpInfo(_fc0)
			fc1 := HideHttpInfo(_fc1)
			fc2 := HideHttpInfo(_fc2)
			fc3 := HideHttpInfo(_fc3)
			fc4 := HideHttpInfo(_fc4)
			fc5 := HideHttpInfo(_fc5)
			fc6 := HideHttpInfo(_fc6)

			errorFunc1 := gin.H{
				"0": _fc0,
				"1": _fc1,
				"2": _fc2,
				"3": _fc3,
				"4": _fc4,
				"5": _fc5,
				"6": _fc6,
			}
			errorHide := gin.H{
				"0": fc0,
				"1": fc1,
				"2": fc2,
				"3": fc3,
				"4": fc4,
				"5": fc5,
				"6": fc6,
			}

			fmt.Println("代码500错误：", errorFunc1)

			// 返回
			ctx.JSON(500, gin.H{
				"state": 500,
				"msg": "代码运行报错，请查看代码运行日志",
				"content": gin.H{
					"gv_version": helper.FrameInfo["gv_version"],
					"error_func": errorHide,
					"error_msg": err,
				},
			})

		}
	}()
	//加载完 defer recover，继续后续接口调用并返回JSON提示
	ctx.Next()
}


