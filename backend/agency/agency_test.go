package main_test

import (
	"testing"

	"github.com/geraldkohn/elian/config"
)

func TestLogin(t *testing.T) {
	err := config.GenerateDB("root:ka112610ng@tcp(127.0.0.1:3306)/test_database?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		t.Error("数据库连接错误: ", err)
	}
}