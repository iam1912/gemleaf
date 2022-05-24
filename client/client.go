package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/iam1912/gemseries/gemleaf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":9091", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	c := pb.NewLeafClient(conn)
	// req := &pb.CreateLeafReq{Name: "test.api.com", Desc: "test.api.com"}
	// _, err = c.CreateLeaf(context.Background(), req)
	// if err != nil {
	// 	fmt.Println(color.RedString("TestLeafServer #%d: create error:%s", 1, err.Error()))
	// }
	req := &pb.GetLeafReq{Name: "test.api.com"}
	resp, err := c.GetLeafID(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}
