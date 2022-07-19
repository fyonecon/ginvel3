package bootstrap

import (
	"ginvel/internal/providers"
	"log"
)

// go级自动运行函数（ < main）
func init() {
	log.Println("init >>> ")
	return
}

// InitChecker 初始化框架、校验框架
// 基础环境检测必须异步阻塞
func (bootstrap *Bootstrap) InitChecker() int64 {
	log.Println("InitChecker >>> ")

	var checker = providers.Checker{}
	var internal = providers.Interval{}

	checker.CheckGolang() // Golang运行环境
	checker.CheckToml(bootstrap.FrameworkConfig) // 检查toml配置文件是否合规
	internal.InitInterval() // 框架心跳，默认周期精度30s

	return 1
}

// InitMustDriver 生成全局服务参数，必选
// 服务任务可选同步、异步
func (bootstrap *Bootstrap) InitMustDriver() int64 {
	log.Println("InitMustProvider >>> ")

	var initMySQL = providers.InitMySQL{}
	var initGORM = providers.InitGORM{}
	var initRedis = providers.InitRedis{}

	// 必选服务----数据库1（默认）
	initMySQL.InitMysql1()
	initGORM.InitGorm1()

	// 必选服务----Redis缓存1
	initRedis.InitRedis1()

	return  1
}

// InitCustomDriver 生成全局服务参数，可选
// 服务任务可选同步、异步
func (bootstrap *Bootstrap) InitCustomDriver() int64 {
	log.Println("InitCustomProvider >>> ")

	var initMySQL = providers.InitMySQL{}
	var initGORM = providers.InitGORM{}
	//var initRedis = providers.InitRedis{}
	//var initES = providers.InitES{}

	// 可选服务----数据库2
	initMySQL.InitMysql2()
	initGORM.InitGorm2()

	// 可选服务----Redis缓存2
	//initMySQL.InitRedis2()

	// 可选服务----搜索和数据分析引
	//initES.InitElasticSearch()

	return  1
}