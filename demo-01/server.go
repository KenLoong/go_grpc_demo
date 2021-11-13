package main

import (
	"net"

	"google.golang.org/grpc"
	"warson.loong/grpc_server_demo/demo-01/services"
)

func main() {
	rpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))

	listener, _ := net.Listen("tcp", ":8888")
	rpcServer.Serve(listener)
}
