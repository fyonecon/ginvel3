package framework_toml

import (
	"log"
	"os"
)

// 空间命名下的全局变量
var frameworkConfig = ReadFrameworkToml()

// GetMySQLConfig1 MySQL数据库配置
func GetMySQLConfig1() map[string]string {

	host := frameworkConfig.Mysql1.Host
	port := frameworkConfig.Mysql1.Port
	name := frameworkConfig.Mysql1.Database
	user := frameworkConfig.Mysql1.User
	pwd := frameworkConfig.Mysql1.Pwd
	charset := frameworkConfig.Mysql1.Encode

	// 默认值
	if len(host) == 0 {
		host = "127.0.0.1"
	}
	if len(port) == 0 {
		port = "3306"
	}
	if len(name) == 0 || len(user) == 0 {
		log.Println("mysql1数据库名或账户名不能为空。。。")
		os.Exit(200)
	}
	if len(charset) == 0 {
		charset = "utf8mb4"
	}

	conf := make(map[string]string)

	conf["DB_HOST"] = host // 127.0.0.1
	conf["DB_PORT"] = port
	conf["DB_NAME"] = name
	conf["DB_USER"] = user
	conf["DB_PWD"] = pwd
	conf["DB_CHARSET"] = charset // "utf8mb4"
	conf["DB_TIMEOUT"] = "12s"

	conf["DB_MAX_OPEN_CONNS"] = "80"       // 连接池最大连接数
	conf["DB_MAX_IDLE_CONNS"] = "10"       // 连接池最大空闲数
	conf["DB_MAX_LIFETIME_CONNS"] = "7200" // 连接池链接最长生命周期

	return conf
}

// GetMySQLConfig2 MySQL数据库配置
func GetMySQLConfig2() map[string]string {

	host := frameworkConfig.Mysql2.Host
	port := frameworkConfig.Mysql2.Port
	name := frameworkConfig.Mysql2.Database
	user := frameworkConfig.Mysql2.User
	pwd := frameworkConfig.Mysql2.Pwd
	charset := frameworkConfig.Mysql2.Encode

	// 默认值
	if len(host) == 0 {
		host = "127.0.0.1"
	}
	if len(port) == 0 {
		port = "3306"
	}
	if len(name) == 0 || len(user) == 0 {
		log.Println("mysql2数据库名或账户名不能为空，但mysql2全局参数将不可用。")
		//os.Exit(200)
	}
	if len(charset) == 0 {
		charset = "utf8mb4"
	}

	conf := make(map[string]string)

	conf["DB_HOST"] = host // 127.0.0.1
	conf["DB_PORT"] = port
	conf["DB_NAME"] = name
	conf["DB_USER"] = user
	conf["DB_PWD"] = pwd
	conf["DB_CHARSET"] = charset // "utf8mb4"
	conf["DB_TIMEOUT"] = "12s"

	conf["DB_MAX_OPEN_CONNS"] = "80"       // 连接池最大连接数
	conf["DB_MAX_IDLE_CONNS"] = "10"       // 连接池最大空闲数
	conf["DB_MAX_LIFETIME_CONNS"] = "7200" // 连接池链接最长生命周期

	return conf
}

// GetRedisConfig1 Redis数据库配置
func GetRedisConfig1() map[string]string {
	// 默认值
	host := frameworkConfig.Redis1.Host
	port := frameworkConfig.Redis1.Port
	pwd := frameworkConfig.Redis1.Pwd
	db := frameworkConfig.Redis1.DataBase
	if len(host) == 0 {
		host = "127.0.0.1"
	}
	if len(port) == 0 {
		port = "6379"
	}
	if len(db) == 0 {
		db = "0"
	}

	conf := make(map[string]string)

	conf["Host"] = host // 例子：127.0.0.1
	conf["Port"] = port // 默认：6379
	conf["Password"] = pwd // 无密码就设置为：""
	conf["DB"] = db // use default DB

	return conf
}

// GetRedisConfig2 Redis数据库配置
func GetRedisConfig2() map[string]string {
	// 默认值
	host := frameworkConfig.Redis2.Host
	port := frameworkConfig.Redis2.Port
	pwd := frameworkConfig.Redis2.Pwd
	db := frameworkConfig.Redis2.DataBase
	if len(host) == 0 {
		host = "127.0.0.1"
	}
	if len(port) == 0 {
		port = "6379"
	}
	if len(db) == 0 {
		db = "0"
	}

	conf := make(map[string]string)

	conf["Host"] = host // 例子：127.0.0.1
	conf["Port"] = port // 默认：6379
	conf["Password"] = pwd // 无密码就设置为：""
	conf["DB"] = db // use default DB

	return conf
}

// GetHttpServerConfig http服务配置
func GetHttpServerConfig() map[string]string {
	host := frameworkConfig.HttpServer.Host
	port := frameworkConfig.HttpServer.Port
	env := frameworkConfig.Common.Env

	// 默认值
	// docker中运行请使用：0.0.0.0，本地测试请使用：127.0.0.1
	if host == "localhost" || len(host) == 0{
		host = "0.0.0.0"
	}
	if len(port) == 0 {
		port = "8090"
	}
	if len(env) == 0 {
		env = "release"
	}

	conf := make(map[string]string)

	conf["Host"] = host // 监听地址，部署在docker中请使用：0.0.0.0。建议不要用127.0.0.1或localhost
	conf["Port"] = port // 监听端口
	conf["Env"] = env // 环境模式 release/debug/test

	return conf
}

// GetViewConfig html模版视图路径配置
func GetViewConfig() map[string]string {
	html := frameworkConfig.View.ViewHTML
	static := frameworkConfig.View.ViewStatic
	tpl := frameworkConfig.View.ViewTPL

	// 默认值
	if len(html) == 0 { // 不需要**/*
		html = "storage/views/html/"
	}
	if len(static) == 0 { // 不需要**/*
		static = "storage/views/static/"
	}
	if len(tpl) == 0 { // 举例："storage/views/html/**/*"
		tpl = "storage/views/html/**/*"
	}

	conf := make(map[string]string)

	// html模板文件路径
	conf["ViewHTML"] = html
	conf["ViewStatic"] = static
	conf["ViewTPL"] = tpl

	return conf
}

//GetElasticSearchConfig ES配置文件
func GetElasticSearchConfig() map[string]string {
	var httpHost string = frameworkConfig.ElasticSearch.HttpHost
	var port string = frameworkConfig.ElasticSearch.Port

	// 必须要有协议头http或https
	if httpHost == "http://localhost" || len(httpHost) == 0{
		httpHost = "http://0.0.0.0"
	}

	if len(port) == 0 {
		port = "9200"
	}

	conf := make(map[string]string)

	conf["HttpHost"] = httpHost
	conf["Port"] = port

	return conf
}