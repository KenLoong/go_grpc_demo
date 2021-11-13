package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"warson.loong/grpc_server_demo/demo-05/server/helper"
	"warson.loong/grpc_server_demo/demo-05/server/services"

	"google.golang.org/grpc"
)

func main() {

	rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCredentials()))
	//注册业务服务，services是生成的pb.go文件所在的包
	//注册商品服务
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	//注册订单服务
	services.RegisterOrderServiceServer(rpcServer, new(services.OrderService))
	//注册用户服务
	services.RegisterUserServiceServer(rpcServer, new(services.UserService))
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("启动网络监听失败 %v\n", err)
	}

	ch := make(chan int, 2)
	go func() {
		//另起协程去启动，因为这个操作会阻塞
		rpcServer.Serve(listen)
		ch <- 1
	}()

	// 使用Gateway启动HTTP Server
	grpcHttpUrl := "localhost:8081" //这个地址是rpc服务的地址
	gwmux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCredentials())}
	//注册商品的gateway服务
	err = services.RegisterProdServiceHandlerFromEndpoint(context.Background(), gwmux, grpcHttpUrl, opt)
	if err != nil {
		log.Fatalf("从GRPC-GateWay连接GRPC失败, err: %v\n", err)
	}
	//注册订单的gateway服务
	err = services.RegisterOrderServiceHandlerFromEndpoint(context.Background(), gwmux, grpcHttpUrl, opt)

	if err != nil {
		log.Fatalf("从GRPC-GateWay连接GRPC失败, err: %v\n", err)
	}

	//对外暴露http接口，从中把请求转换成rpc请求
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}

	go func() {
		//另起协程去启动，因为这个操作会阻塞
		httpServer.ListenAndServe()
		ch <- 1
	}()

	//阻塞主协程
	for i := 0; i < 2; i++ {
		<-ch
	}

}
