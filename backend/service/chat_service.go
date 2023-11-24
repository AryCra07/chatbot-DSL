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

func ChatResponse(userId string, input string) (*pb.ChatResponse, bool) {
	// connect gRPC server
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

	// create gRPC client
	client := pb.NewChatClient(conn)

	// get user info
	user, ok := dao.GetUserInfo(userId)
	if !ok {
		log.Error(consts.Service, "Get user info fail when chat")
		return nil, false
	}

	// prepare request
	request := &pb.ChatRequest{
		State:   user.State,
		Name:    user.Name,
		Input:   input,
		Balance: user.Balance,
		Bill:    user.Bill,
	}

	// call service
	response, err := client.AnswerService(context.Background(), request)
	if err != nil {
		log.Error(consts.Service, "Chat service fail")
		return nil, false
	}
	log.Info(consts.Service, printAnswer(response))

	// update user state
	if user.State != response.State {
		err = dao.UpdateUserState(userId, response.State)
		if err != nil {
			log.Error(consts.Service, "Update user state fail when chat")
			return nil, false
		}
	}

	// update user wallet
	if user.Balance != response.Balance || user.Bill != response.Bill {
		err = dao.UpdateUserWallet(userId, response.Balance, response.Bill)
		if err != nil {
			log.Error(consts.Service, "Update user wallet fail when chat")
			return nil, false
		}
	}

	return response, true
}

// printAnswer print answer
func printAnswer(response *pb.ChatResponse) string {
	if response.Answer == nil {
		return "nil"
	}
	return strings.Join(response.Answer, " <-| ")
}
