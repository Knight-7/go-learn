//Author  :     knight
//Date    :     2020/06/09 15:22:55
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     dota.go

package main

type Data struct {
    Ip       string   `json:"ip"`
    User     string   `json:"user"`
    From     string   `json:"from"`
    Type     string   `json:"type"`
    Content  string   `json:"content"`
    UserList []string `json:"user_list"`
}