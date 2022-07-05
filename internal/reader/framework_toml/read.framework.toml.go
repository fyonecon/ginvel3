package framework_toml

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

var frameworkToml FrameworkConfig

// ReadFrameworkToml 解析toml配置文件，生成全局变量frameworkToml的具体值
// 引用示例：「 var frameworkConfig = read_toml.ReadFrameworkToml() 」
func ReadFrameworkToml() FrameworkConfig {
	// 避免重复解析
	var testHost string = frameworkToml.HttpServer.Host
	if len(testHost) == 0 {
		frameworkToml = LoadFrameworkToml("local") // way = url, local
	}else {
		log.Println("Toml文件已解析，testHost=", testHost)
	}

	return frameworkToml
}


// GetTomlPath 获取toml文件
// way = url, local
func GetTomlPath(way string) (tomlPath string, _way string) {
	// 读取项目绝对路径根目录
	mainDir, _ := os.Getwd()
	mainDirectory := mainDir + "/"

	if way == "url" { // 读取远程文件。请确保远程文件是处于安全的状态，否则会出现账号密码泄露或其他。
		// 读取远程toml文件
		savePath := mainDirectory + "storage/config_toml/" // 读取后将会把文件放在 /storage/toml/
		tomlInfo, _:= DownloadFile(
			"https://makeoss.oss-cn-hangzhou.aliyuncs.com/k8s/config_toml/url.framework.toml",
			savePath,
			"url.framework.toml",
		)
		tomlPath = savePath + ValueInterfaceToString(tomlInfo[0])
	}else if way == "local" {
		// 获取本地toml文件
		tomlPath = mainDirectory + "storage/config_toml/" + "local.framework.toml"
	}else {
		// 默认读storage/config_toml/目录下的文件
		tomlPath = mainDirectory + "storage/config_toml/" + "local.framework.toml"
	}

	return tomlPath, way
}

//LoadFrameworkToml 加载和解析toml文件
func LoadFrameworkToml(way string) (_frameworkConfig FrameworkConfig) {
	// 读取项目绝对路径根目录
	mainDir, _ := os.Getwd()
	mainDirectory := mainDir + "/"

	tomlPath, _way := GetTomlPath(way)

	// 解析toml
	if _, err := toml.DecodeFile(tomlPath, &_frameworkConfig); err != nil {
		Log("项目toml配置读取文件失败", "")
		log.Println("项目toml配置读取文件失败（运行中断）=", _way, err, "")
		os.Exit(200)
	}

	// 动态设置框架根目录
	_frameworkConfig.Framework.FrameworkPath = mainDirectory

	// 设置默认值
	if len(_frameworkConfig.Framework.Timezone) == 0 {
		_frameworkConfig.Framework.Timezone = "Asia/Shanghai"
	}
	if len(_frameworkConfig.Common.Env) == 0 {
		_frameworkConfig.Common.Env = "debug"
	}

	return _frameworkConfig
}
