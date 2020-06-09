//Author  :     knight
//Date    :     2020/06/09 10:00:33
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     tcp/server

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main2() {
	//建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("Dial failed, err:", err)
		return
	}
	//关闭连接
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	//输入信息并发送
	for {
		//读取信息
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}
		//发送信息
		_, err := conn.Write([]byte(inputInfo))
		if err != nil {
			fmt.Println("发送数据失败, err:", err)
			return
		}
		//接收返回的数据
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println("从server接收的数据为:", string(buf[:n]))
	}
}