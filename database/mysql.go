package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

var db *sqlx.DB

//Person 人员信息的实体类
type Person struct {
	UserID   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

//Place 地点的实体类
type Place struct {
	Country string `db:"country"`
	City string `db:"city"`
	TelCode int `db:"telcode"`
}

func init() {
	database, err := sqlx.Open("mysql", "root:@Fight7!@tcp(localhost:3306)/go")

	if err != nil {
		fmt.Println("open mysql failed,error:", err)
		return
	}

	db = database
}

//InsertPerson 插入函数,向数据库中插入数据
func InsertPerson(p *Person) bool {
	r, err := db.Exec("insert into person(username, sex, email) values(?,?,?)", p.Username, p.Sex, p.Email)
	if err != nil {
		fmt.Println("exec failed", err)
		return false
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed", err)
		return false
	}

	fmt.Println("insert sucess", id)
	return true
}

//SelectAll 查找所有的用户
func SelectAll() []Person {
	persons := make([]Person, 0)
	err := db.Select(&persons, "select * from person")
	if err != nil {
		fmt.Println("查找失败")
		return nil
	}
	return persons
}

//Close 关闭数据库连接
func Close() {
	fmt.Println("数据库关闭")
	db.Close()
}

//MysqlTest 测试
func MysqlTest() {
	var op int
	exit := true
	state := `
	1、插入数据
	2、查找所有数据
	3、退出
`
	for exit {
		fmt.Println(state)
		fmt.Println("请选择：")
		fmt.Scanf("%d", &op)
		switch op {
			case 1:
				p := new(Person)
				fmt.Println("请输入姓名，性别， 邮箱：")
				fmt.Scanf("%s %s %s", &p.Username, &p.Sex, &p.Email)
				if InsertPerson(p) {
					fmt.Println("插入成功")
				} else {
					fmt.Println("插入失败")
				}
				time.Sleep(1)
			case 2:
				persons := SelectAll()
				if persons != nil {
					fmt.Println("查找成功")
					for _, v := range persons {
							fmt.Println(v)
						}
				} else {
					fmt.Println("查找失败")
				}
				time.Sleep(1)
			case 3:
				Close()
				exit = false
		}
	}
}