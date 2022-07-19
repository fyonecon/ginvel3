package bootstrap
// Email：fyonecon@gmail.com
// 主页：https://go.ginvel.com
// Apache 2.0
// Ginvel框架并非完全原创，吸纳和引入了很多优秀的开源框架和开发思想，在此感谢各位大佬！Think You！
// 本框架未参考Golang项目结构标准（https://studygolang.com/articles/26941），而是采用了“Laravel+模块化+Vue3+py3”等的思想和实践标准，包含大量的“分层、封装、复用、拆分、迭代、迁移”实践内容。从其他语言或纯go的角度来看待本框架会让你感到不适，建议你多涉语言。

import (
	"ginvel/internal/providers"
	"ginvel/internal/reader/framework_toml"
	_ "go.uber.org/automaxprocs" // 有无docker都可以自动调用Linux多核（win下无法调用，请忽略此问题），详情参考：https://pkg.go.dev/go.uber.org/automaxprocs
	"log"
)

type Bootstrap struct {
	FrameworkConfig framework_toml.FrameworkConfig
}

// RunBase 初始化基础服务
func (bootstrap *Bootstrap) RunBase() {
	log.Println("【运行Base函数】 >>> ")

	bootstrap.FrameworkConfig = framework_toml.ReadFrameworkToml() // 绑定参数

	bootstrap.InitChecker() // 必选
	bootstrap.InitMustDriver() // 必选
	bootstrap.InitCustomDriver() // 可选，自定义

	return
}

// RunServer 初始化服务
func (bootstrap *Bootstrap) RunServer() {
	log.Println("【运行Server函数】 >>> ")

	// 服务停止时清理数据库链接
	defer providers.MysqlDB.Close()
	defer providers.MysqlDB2.Close()

	// 注意顺序
	go bootstrap.RunGrpcServer()  // 必选，启动grpc-server服务
	go bootstrap.TestGrpcClient() // 可选，测试grpc-client服务
	bootstrap.RunHttpServer()     // 必选，启动http-server服务

	return
}
