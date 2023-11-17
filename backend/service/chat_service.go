package service

import (
	"backend/consts"
	"backend/dao"
	"backend/log"
	"backend/pb"
	"context"
	"google.golang.org/grpc"
	"strings"
)

func GetMessage(name string, input string) (*pb.ChatResponse, bool) {
	// 连接 gRPC 服务器
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Error(consts.Service, "gRPC connect fail: %v")
		return nil, false
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Error(consts.Service, "gRPC close fail")
		}
	}(conn)

	// 创建 gRPC 客户端
	client := pb.NewChatClient(conn)

	user, ok := dao.GetUserInfo(name)

	if !ok {
		log.Error(consts.Service, "Get user info fail when chat")
		return nil, false
	}
	// 准备请求
	request := &pb.UserRequest{
		State:   user.State,
		Name:    name,
		Input:   input,
		Balance: user.Balance,
		Bill:    user.Bill,
	}

	// 调用服务
	response, err := client.AnswerService(context.Background(), request)
	if err != nil {
		log.Error(consts.Service, "Chat service fail")
		return nil, false
	}

	log.Info(consts.Service, printAnswer(response))
	err = dao.UpdateUserState(name, response.State)
	if err != nil {
		log.Error(consts.Service, "Update user state fail when chat")
		return nil, false
	}

	err = dao.UpdateUserWallet(name, response.Balance, response.Bill)
	if err != nil {
		log.Error(consts.Service, "Update user wallet fail when chat")
		return nil, false
	}

	return response, true
}

func printAnswer(response *pb.ChatResponse) string {
	if response.Answer == nil {
		return "nil"
	}
	return strings.Join(response.Answer, " <-| ")
}
