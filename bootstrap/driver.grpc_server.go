package bootstrap

import (
	"ginvel/app/request/grpc/grpc_server"
	"ginvel/app/request/grpc/proto"
	"ginvel/internal/middlewares"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery" // 服务重启
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

// RunGrpcServer 启动grpc服务端，Powered by gRPC。
func (bootstrap *Bootstrap) RunGrpcServer() {
	var host string = bootstrap.FrameworkConfig.GRPCServer.Host
	var port string = bootstrap.FrameworkConfig.GRPCServer.Port
	var address string = host + ":" + port

	// 监听
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Println("grpc-server服务监听失败: ", err)
		os.Exit(200)
	}

	// 实例化grpc服务端，并插入中间价
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer( // 流式拦截器
			grpcRecovery.StreamServerInterceptor(),
			middlewares.StreamGSError500(address),

		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer( // 简单拦截器
			grpcRecovery.UnaryServerInterceptor(),
			middlewares.UnaryGSError500(address),

		)),
	)

	// 注册Greeter服务，监听目标函数
	proto.RegisterGreeterServer(grpcServer, &grpc_server.GrpcServer{})

	// 往grpc服务端注册反射服务
	reflection.Register(grpcServer)

	log.Println("grpc-server访问地址示例：http://" + address + " >>> \n ")

	if err := grpcServer.Serve(lis); err != nil {
		log.Println("grpc-server服务连接失败: ", err)
		os.Exit(200)
	}

	return
}