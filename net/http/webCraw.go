//Author  :     knight
//Date    :     2020/06/09 14:54:30
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     试试看go的网络爬虫

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.topgoer.com/go%E5%9F%BA%E7%A1%80/Go%E8%AF%AD%E8%A8%80%E7%9A%84%E4%B8%BB%E8%A6%81%E7%89%B9%E5%BE%81.html")
	if err != nil {
		fmt.Println("获取网络资源失败, err:", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("body: %T\n,value: %v", body, string(body))
	
}