package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"warson.loong/grpc_server_demo/demo-02/client/services"
)

func main() {

	creds, err := credentials.NewClientTLSFromFile("keys/server.crt", "warson.com")
	if err != nil {
		log.Fatalf("加载客户端证书失败, err: %v\n", err)
	}

	//GODEBUG 为 x509ignoreCN=0
	//set GODEBUG=x509ignoreCN=0
	//但是需要注意，从 Go 1.17 开始，环境变量就不再生效了，必须通过 SAN 方式才行。所以，为了后续的 Go 版本升级，还是早日支持为好。
	conn, err := grpc.Dial(":8888", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("连接GRPC服务端失败 %v\n", err)
	}
	defer conn.Close()

	prodClient := services.NewProdServiceClient(conn)

	prodResp, err := prodClient.GetProdStock(context.Background(), &services.ProdRequest{ProdId: 12})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(prodResp.ProdStock)
}
