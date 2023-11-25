package service

import (
	"backend/consts"
	"backend/dao"
	"backend/log"
	"backend/pb"
	"context"
	"google.golang.org/grpc"
)

// Timer implement timer service
/*
 * @param userId: user's id
 * @return []string: answer
 * @return bool: whether the user is exit
 * @return bool: whether the user is reset
 * @return bool: whether the service is fail
 */
func Timer(userId string, lastTime int32, nowTime int32) (*pb.TimerResponse, bool) {
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
	client := pb.NewTimerClient(conn)

	user, ok := dao.GetUserInfo(userId)

	if !ok {
		log.Error(consts.Service, "Timer service fail: user does not exist")
		return nil, false
	}

	// 准备请求
	request := &pb.TimerRequest{
		State:    user.State,
		NowTime:  nowTime,
		LastTime: lastTime,
	}

	// 调用服务
	response, err := client.TimerService(context.Background(), request)
	if err != nil {
		log.Error(consts.Service, "Timer service fail: %v")
		return nil, false
	}

	//log.Info(consts.Service, printAnswer(response))
	if user.State != response.State {
		err = dao.UpdateUserState(userId, response.State)
		if err != nil {
			log.Error(consts.Service, "Update user state fail when timer")
			return nil, false
		}
	}

	return response, true
}
