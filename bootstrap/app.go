package bootstrap
// Email：fyonecon@gmail.com
// 主页：https://go.ginvel.com
// Apache 2.0
// Ginvel框架并非完全原创，吸纳和引入了很多优秀的开源框架和开发思想，在此感谢各位大佬！Think You！

import (
	"ginvel/internal/providers"
	"ginvel/internal/reader/framework_toml"
	_ "go.uber.org/automaxprocs" // 有无docker都可以自动调用Linux多核（win下无法调用，请忽略此问题），详情参考：https://pkg.go.dev/go.uber.org/automaxprocs
	"log"
)

type Bootstrap struct {
	FrameworkConfig framework_toml.FrameworkConfig
}

// 框架运行环境自检
func init() {
	var checker = providers.Checker{}
	checker.CheckGolang()
	return
}

// InitApp 初始化框架、校验框架、生成全局参数
func (bootstrap *Bootstrap) InitApp()  {
	log.Println("【InitServer】 >>> ")

	var checker = providers.Checker{}
	var initMySQL = providers.InitMySQL{}
	var initGORM = providers.InitGORM{}
	var initInternal = providers.InitInterval{}
	var initRedis = providers.InitRedis{}
	var initES = providers.InitES{}

	// 必选组件-----------------------
	// 检查toml配置文件是否合规
	checker.CheckToml(bootstrap.FrameworkConfig)
	// 数据库1（默认）
	initMySQL.InitMysql1()
	initGORM.InitGorm1()
	// 数据库2
	//initMySQL.InitMysql2()
	//initGORM.InitGorm2()

	// 框架心跳，默认周期精度30s
	initInternal.InitInterval()

	// 可选组件-----------------------
	// Redis缓存1
	initRedis.InitRedis1()
	// Redis缓存2
	//initMySQL.InitRedis2()
	// 搜索和数据分析引
	initES.InitElasticSearch()

	return
}

// RunApp Ginvel框架服务启动（入口）
// 本框架未参考Golang项目结构标准（https://studygolang.com/articles/26941），而是采用了“Laravel+模块化+Vue3+py3”等的思想和实践标准，包含大量的“分层、封装、复用、拆分、迭代、迁移”实践内容。从其他语言或纯go的角度来看待本框架会让你感到不适，建议你多涉语言。
func (bootstrap *Bootstrap) RunApp() {
	log.Println("【运行App函数】 >>> ")

	// 初始化基础服务
	bootstrap.FrameworkConfig = framework_toml.ReadFrameworkToml()
	bootstrap.InitApp()

	// 服务停止时清理数据库链接
	defer providers.MysqlDB.Close()
	defer providers.MysqlDB2.Close()

	// 注意顺序
	go bootstrap.RunGrpcServer()  // 必选，启动grpc-server服务
	go bootstrap.TestGrpcClient() // 可选，测试grpc-client服务
	bootstrap.RunHttpServer()     // 必选，启动http-server服务

	return
}
