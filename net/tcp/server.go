//Author  :     knight
//Date    :     2020/06/09 09:51:19
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     tcp/client

package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		//接受数据
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			return
		}
		recvStr := string(buf[:n])
		fmt.Println("接受到client端发来的数据：", recvStr)

		//将数据发送回去
		conn.Write([]byte(recvStr))
	}
}

func main1() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err", err)
			continue
		}
		go process(conn)
	}
}