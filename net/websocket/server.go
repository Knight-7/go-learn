//Author  :     knight
//Date    :     2020/06/09 15:12:01
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     一个websocket的服务端程序

package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
    go h.run()
    router.HandleFunc("/ws", myws)
    if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
        fmt.Println("err:", err)
    }
}