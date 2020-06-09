//Author  :     knight
//Date    :     2020/06/09 14:41:41
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     http/server

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main2() {
	resp, _ := http.Get("http://localhost:2020/go")
	fmt.Println(resp)
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header)

	buf := make([]byte, 1024)
	for {
		//接收服务端的信息
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
		} else {
			fmt.Println("读取完毕")
			res := string(buf[:n])
			fmt.Println(res)
			break
		}
	}
}