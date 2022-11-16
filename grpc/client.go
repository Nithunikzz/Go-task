package main

import (
	"context"
	"log"

	chat "github.com/connect2naga/go-examples/grpc/unary/proto"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dail the port 50052, error: %v ", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	res, err := c.SayHello(context.Background(), &chat.MessageReq{Body: "Hi this is nithu...."})
	if err != nil {
		log.Fatalf("error when calling sayHello, error:%v", err)
	}
	log.Printf("response from server : %s", res.Body)
}
