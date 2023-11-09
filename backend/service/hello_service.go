package service

import (
	"backend/consts"
	"backend/log"
	"backend/pb"
	"context"
	"google.golang.org/grpc"
)

func GetHello(name string, state int32, wallet map[string]int32) []string {
	// 连接 gRPC 服务器
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Error(consts.Service, "gRPC connect fail: %v")
		return nil
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Error(consts.Service, "gRPC close fail")
		}
	}(conn)

	// 创建 gRPC 客户端
	client := pb.NewGreetClient(conn)

	// 准备请求
	request := &pb.UserRequest{
		State:  state,
		Name:   name,
		Wallet: wallet,
	}

	// 调用服务
	response, err := client.SayHelloService(context.Background(), request)
	if err != nil {
		log.Error(consts.Service, "Hello service fail")
		return nil
	}

	return response.Words
}
