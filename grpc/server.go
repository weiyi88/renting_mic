package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	pd "renting_mic/myproto"
)

type server struct {
	pd.UnsafeHelloserverServer
}

func (this *server) Sayhello(context.Context, *pd.HelloReq) (*pd.HelloRsp, error) {
	return &pd.HelloRsp{Msg: "hello bubibibib"}, nil
	//return nil, status.Errorf(codes.Unimplemented, "method Sayhello not implemented")
}
func (this *server) Sayname(context.Context, *pd.NameReq) (*pd.NameRsp, error) {
	return &pd.NameRsp{Msg: "fuckfuckfuck"}, nil
	//return nil, status.Errorf(codes.Unimplemented, "method Sayname not implemented")
}
func main() {

	// 使用tcp 监听网络
	ln, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("网络错误：", err)
	}

	// 创建grpc服务对象
	srv := grpc.NewServer()

	// 注册服务
	pd.RegisterHelloserverServer(srv, &server{})

	// 等待网络链接
	err = srv.Serve(ln)
	if err != nil {
		fmt.Println("网络错误", err)
	}
}
