package middlewares
// gRPC-Client只作为主动调用接口没必要去做中间件，所以这里仅以gRPC-Server演示

import (
	"context"
	"fmt"
	"ginvel/common/helper"
	"ginvel/common/kits"
	"google.golang.org/grpc"
)


// StreamGSError500 捕捉流式代码致命错误
func StreamGSError500(address string) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {

		// 自定义中间件代码
		fmt.Println("StreamGSError500 服务已经加入监听===")
		defer func() {
			if err := recover(); err != nil {
				//打印错误堆栈信息
				fmt.Println(err)
				// debug.PrintStack() // 显示报错详情
				kits.Error(helper.ValueInterfaceToString(err), address)
			}
		}()

		// 固定返回
		err = handler(srv, stream)
		return err
	}
}

// UnaryGSError500 捕捉简单代码致命错误
func UnaryGSError500(address string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {

		// 自定义中间件代码
		fmt.Println("UnaryGSError500 服务已经加入监听===")
		defer func() {
			if err := recover(); err != nil {
				//打印错误堆栈信息
				fmt.Println(err)
				// debug.PrintStack() // 显示报错详情
				kits.Error(helper.ValueInterfaceToString(err), address)
			}
		}()

		// 固定返回
		resp, err := handler(ctx, req)
		return resp, err
	}
}
