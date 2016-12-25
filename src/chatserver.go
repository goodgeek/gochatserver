package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Hello world!!")

	listener, err := net.Listen("tcp4", "8002")
	if err != nil {
		fmt.Println("启动服务器失败:" + err.Error())
		return
	}

	fmt.Println("启动服务器成功 端口:" + "8002")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept 失败:" + err.Error())
		}

		conn.Close()
	}
}
