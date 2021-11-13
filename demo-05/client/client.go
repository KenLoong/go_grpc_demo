package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"warson.loong/grpc_server_demo/demo-05/server/helper"
	"warson.loong/grpc_server_demo/demo-05/server/services"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(helper.GetClientCredentials()))
	if err != nil {
		log.Fatalf("连接GRPC服务端失败 %v\n", err)
	}

	defer conn.Close()

	// prodClient := services.NewProdServiceClient(conn)
	//请求切片
	// prodRes, err := prodClient.GetProdStocks(context.Background(),
	// 	&services.QuerySize{Size: 5})

	// if err != nil {
	// 	log.Fatalf("请求GRPC服务端失败 %v\n", err)
	// }
	// fmt.Println(prodRes.Prodres)

	// rsp, err := prodClient.GetProdStock(context.Background(), //请求商品库存
	// 	&services.ProdRequest{ProdId: 2, ProdArea: 1})
	// if err != nil {
	// 	log.Fatalf("请求GRPC服务端失败 %v\n", err)
	// }

	// resp, err := prodClient.GetProdInfo(context.Background(),
	// 	&services.ProdRequest{ProdId: 1}) //请求商品信息

	/*
		OrderId    int32                  `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"` //订单id，数字自增
		OrderNo    string                 `protobuf:"bytes,2,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
		UserId     int32                  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
		OrderMoney float32                `protobuf:"fixed32,4,opt,name=order_money,json=orderMoney,proto3" json:"order_money,omitempty"` //商品金额
		OrderTime
	*/

	// ti := timestamp.Timestamp{Seconds: time.Now().Unix()} //google第三方包
	// fmt.Println(ti)
	// orderClient := services.NewOrderServiceClient(conn)
	// resp, err := orderClient.NewOrder(context.Background(),
	// 	&services.OrderMain{
	// 		OrderId:    2,
	// 		OrderNo:    "s5",
	// 		UserId:     4,
	// 		OrderMoney: 11.3,
	// 		OrderTime:  &ti,
	// 	})
	// fmt.Println(resp)

	//请求服务端流
	// userClient := services.NewUserServiceClient(conn)
	// var i int32
	// req := services.UserScoreRequest{}
	// req.Users = make([]*services.UserInfo, 0)
	// for i = 1; i < 8; i++ {
	// 	req.Users = append(req.Users, &services.UserInfo{UserId: i})
	// }

	// res, err := userClient.GetUserScore(context.Background(), &req)
	// fmt.Println(res.Users)

	// stream, err := userClient.GetUserScoreServerStream(context.Background(), &req)
	// if err != nil {
	// 	fmt.Println("GetUserScoreServerStream err:", err.Error())
	// 	return
	// }

	// for {
	// 	//接收服务端响应
	// 	res, err := stream.Recv()
	// 	if err == io.EOF {
	// 		break //读完了
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	//每次都打印出2条信息
	// 	fmt.Println(res.Users)
	// }

	// //客户端流
	// userClient := services.NewUserServiceClient(conn)
	// stream, err := userClient.GetUserScoreClientStream(context.Background())
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// for i := 1; i <= 3; i++ {
	// 	req := new(services.UserScoreRequest)
	// 	req.Users = make([]*services.UserInfo, 0)
	// 	for j := 1; j <= 3; j++ {
	// 		//假设过程耗时
	// 		req.Users = append(req.Users, &services.UserInfo{UserId: int32(j)})
	// 	}

	// 	err = stream.Send(req)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// }
	// resp, err := stream.CloseAndRecv()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(resp.Users)

	userClient := services.NewUserServiceClient(conn)
	stream, err := userClient.GetUserScoreTwsStream(context.Background())

	if err != nil {
		log.Fatalf("请求GRPC服务端失败 %v\n", err)
	}

	for i := 0; i < 3; i++ {
		req := new(services.UserScoreRequest)
		req.Users = make([]*services.UserInfo, 0)
		var j int32
		for j = 1; j <= 5; j++ {
			req.Users = append(req.Users, &services.UserInfo{UserId: j})
		}
		//发送
		stream.Send(req)

		//接收
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("接收服务端请求失败 %v\n", err)
		}

		fmt.Println(res.Users)

	}
}
