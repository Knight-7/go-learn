//Author  :     knight
//Date    :     2020/06/09 10:40:44
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     http/server

package main

import (
	"fmt"
	"net/http"
)

func main1() {
	//http:127.0.0.1:3434/go
	// 单独写回调函数
	http.HandleFunc("/go", myHandler)
	//http.HandleFunc("/ungo", myHandler)
	//addr: 监听的地址
	//handler: 回调函数
	err := http.ListenAndServe("127.0.0.1:2020", nil)
	if err != nil {
		fmt.Println("listen and server failed, err:", err)
	}
}

//handler函数
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	//请求方式： GET,POST,DELETE,PUT,UPDATE
	fmt.Println("method:", r.Method)
	// /go
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	// 回复
	w.Write([]byte("I'm Knight."))
}