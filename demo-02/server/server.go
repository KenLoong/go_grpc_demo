package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"warson.loong/grpc_server_demo/demo-02/server/services"
)

func main() {
	creds, err := credentials.NewServerTLSFromFile("keys/server.crt", "keys/server_no_passwd.key")
	if err != nil {
		log.Fatalf("加载服务端证书和Key失败, err: %v\n", err)
	}

	rpcServer := grpc.NewServer(grpc.Creds(creds))
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("启动网络监听失败 %v\n", err)
	}
	rpcServer.Serve(listen)
}
