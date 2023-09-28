package main

import (
	"context"
	"fmt"
	"log"

	"backend/pb" // 替换为你的 Protobuf 包的实际导入路径
	"google.golang.org/grpc"
)

func main() {
	// 连接 gRPC 服务器
	conn, err := grpc.Dial("127.0.0.1:8972", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()

	// 创建 gRPC 客户端
	client := pb.NewGreeterClient(conn)

	// 创建上下文和请求
	ctx := context.Background()
	req := &pb.HelloRequest{Name: "AryCra07"}

	// 调用 gRPC 服务
	resp, err := client.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("调用服务失败: %v", err)
	}

	// 打印响应
	fmt.Printf("服务响应: %s\n", resp.GetReply())
}
