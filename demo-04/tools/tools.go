package tools

// https://github.com/grpc-ecosystem/grpc-gateway

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)

//两个都执行一次
// import (
// 	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
// 	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
// 	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
// )
