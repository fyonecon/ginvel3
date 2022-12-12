package main

import (
	"ginvel/bootstrap"
)

// @title Ginvel微服务框架
// @version gv-3.+
// @description Ginvel是一个MVC+RPC的Golang微服务框架，Http基于Gin，DDD服务基于gRPC。
// @contact.email fyonecon@gmail.com
// @contact.url https://go.ginvel.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	var run = bootstrap.Bootstrap{}
	run.RunBase()
	run.RunServer()
}
