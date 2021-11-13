package main

import (
	"context"
	"fmt"
	"log"

	"warson.loong/grpc_server_demo/demo-04/server/helper"
	"warson.loong/grpc_server_demo/demo-04/server/services"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(helper.GetClientCredentials()))
	if err != nil {
		log.Fatalf("连接GRPC服务端失败 %v\n", err)
	}

	defer conn.Close()
	prodClient := services.NewProdServiceClient(conn)
	prodRes, err := prodClient.GetProdStock(context.Background(),
		&services.ProdRequest{ProdId: 12})

	if err != nil {
		log.Fatalf("请求GRPC服务端失败 %v\n", err)
	}
	fmt.Println(prodRes.ProdStock)
}
