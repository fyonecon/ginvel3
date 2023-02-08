package grpc_server

import (
	"context"
	"fmt"
	"ginvel/app/request/grpc/grpc_server/GSControllers/GSExample"
	"ginvel/app/request/grpc/proto"
)

// GrpcServer 用来实现proto文件里面的Greeter服务里面的接口
type GrpcServer struct {
	proto.UnimplementedGreeterServer
}

// GrpcHandler rpc目标调用函数
// 函数及其传参映射，方便Server-Client之间调用
func (s *GrpcServer) GrpcHandler (ctx context.Context, in *proto.HandlerInput) (*proto.HandlerBack, error) {
	var inputMethod string = in.GetInputMethod() // 或用in.InputMethod
	var inputData string = in.GetInputData()
	var inputToken string = in.GetInputToken()

	var backState int64 = -1
	var backMsg string = ""
	var backData string = ""
	var backToken string = "back-token-2021"
	var backTest string = ""

	// 验证inputToken
	tokenState, tokenMsg := GSCheckToken(inputToken)
	if tokenState == 1 { // token正确

		// 方法函数映射
		switch inputMethod {
		case "FuncTest1":
			backState, backMsg, backData = GSExample.FuncTest1(inputData)
			break
		case "FuncTest2":
			backState, backMsg, backData = GSExample.FuncTest2(inputData)
			break
		default: // 无对应映射
			fmt.Println("GrpcHandler无默认input-method：", inputMethod, inputData)
			backState = 0
			backMsg = "GrpcHandler无默认input-method"
			backData = ""
			break
		}

	}else {
		backState = 0
		backMsg = tokenMsg
		backData = ""
	}

	// 处理完参数，给client返回值
	return &proto.HandlerBack{
		BackState: backState,
		BackMsg: backMsg,
		BackData: backData,
		BackToken: backToken,
		BackTest: backTest,
	}, nil
}