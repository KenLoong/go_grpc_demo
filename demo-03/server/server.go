package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc/credentials"
	"warson.loong/grpc_server_demo/demo-03/server/services"

	"google.golang.org/grpc"
)

func main() {

	// 证书认证-双向认证
	// 从证书相关文件中读取和解析信息，得到证书公钥、密钥对
	cert, err := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	if err != nil {
		log.Fatalf("加载服务端证书失败, err: %v\n", err)
	}

	certPool := x509.NewCertPool() //创建证书池
	ca, err := ioutil.ReadFile("cert/ca.pem")
	if err != nil {
		log.Fatalf("读取公钥文件失败: %v\n", err)
	}
	// 尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
	certPool.AppendCertsFromPEM(ca) //把证书放进证书池

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},        // 设置证书链，允许包含一个或多个
		ClientAuth:   tls.RequireAndVerifyClientCert, // 要求必须校验客户端的证书。可以根据实际情况选用以下参数
		ClientCAs:    certPool,                       // 设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
	})

	rpcServer := grpc.NewServer(grpc.Creds(creds))
	//注册业务服务，services是生成的pb.go文件所在的包
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("启动网络监听失败 %v\n", err)
	}
	rpcServer.Serve(listen)
}
