package main

import (
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
	"github.com/zhoushuguang/lebron/cmd/user/rpc/model"
)

func main() {
	userFieldNames := builder.RawFieldNames(&model.User{})
	userRowsExpectAutoSet := strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	userRowsWithPlaceHolder := strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
	fmt.Println(userFieldNames, userRowsExpectAutoSet, userRowsWithPlaceHolder)
}
