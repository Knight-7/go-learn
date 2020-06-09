package test

import (
	"testing"
)

func TestAdd(t * testing.T) {
	got := Add(7, 2, 3, 4)
	want := 16
	if got != want {
		t.Errorf("测试失败,预期结果为：%d, 函数结果为：%d", want, got)
	}
}