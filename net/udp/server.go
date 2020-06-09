//Author  :     knight
//Date    :     2020/06/09 10:14:57
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     udp/server

package main

import (
	"fmt"
	"net"
)

func main1() {
	//开始连接
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP: net.IPv4(0, 0, 0, 0),
		Port: 20000,
	})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	//关闭连接
	defer listen.Close()
	//接收数据
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read udp failed, err:", err)
			continue
		}
		fmt.Printf("data: %v, addr: %v, count: %v\n", string(data[:n]), addr, n)
		//发送数据
		_, err = listen.WriteToUDP(data[:n], addr)
		if err != nil{
			fmt.Println("write to udp failed, err:", err)
			continue 
		}
	}
}