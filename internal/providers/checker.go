package providers

import (
	"ginvel/common/helper"
	"ginvel/internal/reader/framework_toml"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type Checker struct {}

// CheckGolang 检查Go环境
// 默认64位
func (checker *Checker) CheckGolang()  {
	var _goVersion string = runtime.Version()
	var minVersion string = "go1.17.0"
	var goVersion string = strings.Replace(_goVersion, "go", "", -1)
	var goArray []string = strings.Split(goVersion, ".")
	var go1 int
	var go2 int
	var goYes int
	for _, value := range goArray{
		theValue, _ := strconv.ParseInt(value, 10, 64)
		if go1 == 0 {
			go1 = int(theValue)
		}else {
			if go2==0 {
				go2 = int(theValue)
				if go1 == 1 { // min
					goYes = goYes + 1
					if go2 >= 17 { // min
						goYes = goYes + 1
					}
				}else if go1 >= 2 {
					goYes = goYes + 2
				}
			}else {
				break
			}
		}
	}
	if goYes == 2 {
		//log.Println("Go版本符合要求 >>> 当前Go版本：", _goVersion, " 最小要求版本：" + minVersion, " Golang权威文档：https://go.dev/ 或 https://golang.google.cn/ ")
	}else {
		log.Println("Go版本太低，框架所依赖的特性将不可直接使用！Ginvel服务自动中断！", "当前Go版本："+_goVersion, " 最小要求版本：" + minVersion, "运行中断")
		os.Exit(200)
	}

	return
}

// CheckToml 检测Toml是否与当前框架匹配
func (checker *Checker) CheckToml(frameworkConfig framework_toml.FrameworkConfig)  {
	var falseFrameworkHashToml string = frameworkConfig.Framework.FrameworkHash

	// 自杀程序，需要toml hash的格式：year@密钥 ，code_config格式：密钥
	var _hashYear int64 = 25 // 最后年份Y（界限年份，前后有效），如：24是2024年
	var _year int64 = _hashYear
	var _frameworkHashToml string = ""
	var hashArray []string = strings.Split(falseFrameworkHashToml, "@")
	if len(hashArray) >= 2 {
		_hashYear = helper.StringToInt(hashArray[0])
		_frameworkHashToml = hashArray[1]
	}else {
		_hashYear = _year + 2
		_frameworkHashToml = falseFrameworkHashToml
	}
	if _hashYear > _year + 5 { _hashYear = _year + 2 }else if _hashYear < _year + 1  {_hashYear = _year + 1} // 指在不维护的情况下，为了系统与数据的安全，可以正常运行[1--5年]。仅限在新启动程序的情况下，对已经运行对程序无影响。
	var nowYear int64 = helper.StringToInt(helper.GetTimeDate("Y")) - 2000
	if nowYear > _hashYear { // 不可用
		//log.Println("安全检测不通过，触发自杀程序。")
		os.Exit(200)
	}

	// 安全检测
	var frameworkHashToml string = helper.MD5(_frameworkHashToml) + "+hash"
	var frameworkHashGo string =  helper.FrameInfo["framework_hash"]
	if frameworkHashToml == frameworkHashGo {
		log.Println("framework_hash符合要求 >>> ", "framework_hash="+_frameworkHashToml)
	}else{
		log.Println("framework_hash不匹配，toml文件解析失败，Ginvel服务自动中断！", "framework_hash="+_frameworkHashToml)
		os.Exit(200)
	}

	return
}
