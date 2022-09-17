package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
)

func pandatext(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello hello Aring")
}

type Panda int

// 标准格式， 函数关键字（对象）函数名（对端发送过来的内容，返回给对端的内容）错误
func (this *Panda) Getinfo(argType int, replyType *int) error {
	fmt.Println("打印对端发送过来的内容为：", argType)

	// 修改内容之
	*replyType = argType + 100
	return nil
}

func main() {
	http.HandleFunc("/panda", pandatext)

	// 不new 指针没内存
	pd := new(Panda)
	// 服务端注册一个对象，暴露出来
	rpc.Register(pd)
	rpc.HandleHTTP()

	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Print("网络错误")
	}
	http.Serve(ln, nil)

}
