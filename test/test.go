package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"renting_mic/prototest"
)

func main() {
	test := &prototest.Test{
		Name:    "Aring",
		Stature: 23,
		Weight:  []int64{120, 125},
		Motto:   "HELLO",
	}

	fmt.Println(test)

	// proto 编码
	data, err := proto.Marshal(test)
	if err != nil {
		fmt.Println("编码失败")
	}
	fmt.Println(data)
}
