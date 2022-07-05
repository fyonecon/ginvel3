package GCExample

import (
	"context"
	"ginvel/app/request/grpc/grpc_client"
	"ginvel/app/request/grpc/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

// TestGrpc1 grpc客户端访问示例
// 需要实时调用Dial()
func TestGrpc1(host string, port string) map[string]interface{} {
	address := host + ":" + port
	log.Println("grpc-client >>> ", address)

	defer func() {
		if r := recover(); r != nil {
			log.Println("grpc-client出现致命错误，已经跳过该错误。")
		}
	}()

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("不能连接grpc-client服务: %v", err)
	}
	defer conn.Close()
	grpcClient := proto.NewGreeterClient(conn)

	//
	grpcClientContext, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//
	// client调用server函数
	function, err := grpcClient.GrpcHandler(grpcClientContext, &proto.HandlerInput{
		InputMethod: "FuncTest1", // 方法函数映射
		InputData: "id=1&year=2021", // 参数
		InputToken: "input-token-2021", // token
		InputTest: "id=2&year=2030", // 测试数据
	})

	var backState int64 = 0
	var backMsg string
	var backToken string
	var backData string
	var backTest string

	if err != nil {
		log.Fatalf("不能连接grpcClient: %v", err)
		backMsg = "不能连接grpcClient"
	}else {

		// 从server返回的值
		backState = function.GetBackState()
		backMsg = function.GetBackMsg()
		backToken = function.GetBackToken()
		backData = function.GetBackData()
		backTest = function.GetBackTest()

		// 验证backToken
		tokenState, tokenMsg := grpc_client.GCCheckToken(backToken)
		if tokenState != 1 { // token不正确
			log.Println("已连接grpcClient，token不正确：", tokenState, tokenMsg)
			backState = 0
			backMsg = "已连接grpcClient，token不正确"
			backData = ""
		}else { // token正确
			// 其他操作

		}

	}

	var back = map[string]interface{}{
		"backState": backState,
		"backMsg": backMsg,
		"backData": backData,
		"backTest": backTest,
	}

	return back
}
