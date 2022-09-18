package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pd "renting_mic/myproto"
)

func main() {

	// 链接服务端
	connect, err := grpc.Dial("127.0.0.1:8001", grpc.WithInsecure()) // 固定参数固定写法
	if err != nil {
		fmt.Println("网络异常：", err)
	}

	// 网络延迟关闭
	defer connect.Close()

	// 获得grpc 句柄
	c := pd.NewHelloserverClient(connect)

	//通过句柄，调用函数
	rehello, err := c.Sayhello(context.Background(), &pd.HelloReq{Name: "呼哈"})
	if err != nil {
		fmt.Println("调用sayHello失败")
	}

	fmt.Println("调用sayHello 返回：", rehello.Msg)

	rename, err := c.Sayname(context.Background(), &pd.NameReq{Name: "holdor"})
	if err != nil {
		fmt.Println("调用Sayname 失败：", err)
	}
	fmt.Println("调用Sayname 返回：", rename.Msg)
}
