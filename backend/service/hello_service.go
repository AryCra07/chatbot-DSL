package service

import (
	"backend/consts"
	"backend/dao"
	"backend/log"
	"backend/pb"
	"context"
	"google.golang.org/grpc"
)

func Hello(userId string) []string {
	// connect gRPC server
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

	// create gRPC client
	client := pb.NewGreetClient(conn)

	user, ok := dao.GetUserInfo(userId)

	if !ok {
		log.Error(consts.Service, "Hello service fail")
		return nil
	}

	// prepare request
	request := &pb.HelloRequest{
		State: user.State,
		Name:  user.Name,
	}

	// call service
	response, err := client.SayHelloService(context.Background(), request)
	if err != nil {
		log.Error(consts.Service, "Hello service fail")
		return nil
	}

	return response.Words
}
