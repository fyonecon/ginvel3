package bootstrap

import (
	"ginvel/app/request/grpc/grpc_client/GCControllers/GCExample"
	"log"
)

// TestGrpcClient 启动测试的grpc客户端，Powered by gRPC。
// 测试grpc-client功能，client连接server需要每个client请求中每次dial
func (bootstrap *Bootstrap) TestGrpcClient() map[string]interface{}{
	var host string = bootstrap.FrameworkConfig.GRPCClient.Host
	var port string = bootstrap.FrameworkConfig.GRPCClient.Port

	var back = GCExample.TestGrpc1(host, port)
	log.Println("测试GrpcClient连接情况 >>> ", back)
	return back
}