package main

import (
	"fmt"
	"net/rpc"
)

func main() {

	// 建立网络链接
	cli, err := rpc.DialHTTP("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("网络链接错误")
	}

	var pd int

	// 调用rpc 服务端方法
	err = cli.Call("Panda.Getinfo", 101, &pd)
	if err != nil {
		fmt.Println("打call 失败")
	}
	fmt.Println("最后得到值为：", pd)

}
