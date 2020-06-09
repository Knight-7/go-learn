//Author  :     knight
//Date    :     2020/06/09 10:23:27
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     udp/client

package main

import (
	"fmt"
	"net"
)

func main2() {
	//连接服务器
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP: net.IPv4(0, 0, 0, 0),
		Port: 20000,
	})
	if err != nil {
		fmt.Println("连接服务器失败")
		return
	}
	//关闭连接
	defer socket.Close()
	//发送数据
	sendData := []byte("Hello, server. I'm Knight.")
	_, err = socket.Write(sendData)
	if err != nil {
		fmt.Println("发送数据失败")
		return
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println("接收数据失败, err:", err)
		return
	}
	fmt.Printf("recv: %v, addr: %v, count: %v\n", string(data[:n]), remoteAddr, n)
}