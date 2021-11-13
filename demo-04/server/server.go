package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"warson.loong/grpc_server_demo/demo-04/server/helper"
	"warson.loong/grpc_server_demo/demo-04/server/services"

	"google.golang.org/grpc"
)

func main() {

	rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCredentials()))
	//注册业务服务，services是生成的pb.go文件所在的包
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
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
	gwmux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCredentials())}
	err = services.RegisterProdServiceHandlerFromEndpoint(context.Background(), gwmux, "localhost:8081", opt)

	if err != nil {
		log.Fatalf("从GRPC-GateWay连接GRPC失败, err: %v\n", err)
	}
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
