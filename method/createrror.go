package main

import (
	"fmt"
	"os"
	"time"
)

//PathError 是一个自定义的路径error
type PathError struct {
	path string
	op string
	createTime string
	message string
}

func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s\n op=%s\n createTime=%s\n message=%s", p.path, p.op, p.createTime, p.message)
}

//Open open a file
func Open(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return  &PathError{
			path:        fileName,
			op:          "read",
			message:     err.Error(),
			createTime:  fmt.Sprintf("%v", time.Now()),
		}
	}

	defer file.Close()
	return nil
}

func main1() {
	err := Open("/Users/Documents/test.txt")
	switch v := err.(type) {
	case *PathError:
		fmt.Println("get path error", v)
	default:
	}

}