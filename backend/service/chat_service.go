package service

import (
	"backend/consts"
	"backend/log"
	"backend/pb"
	"context"
	"google.golang.org/grpc"
)

func GetMessage(name string, state int32, input string) []string {
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
	client := pb.NewChatClient(conn)

	// 准备请求
	request := &pb.UserRequest{
		State:  state,
		Name:   name,
		Input:  input,
		Wallet: map[string]int32{},
	}

	// 调用服务
	response, err := client.AnswerService(context.Background(), request)
	if err != nil {
		log.Error(consts.Service, "Chat service fail")
		return nil
	}

	return response.Answer
}
