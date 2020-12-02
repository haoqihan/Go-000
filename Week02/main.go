package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"week02/internal/service"
)

func main() {
	svc := service.New()
	userlist, err := svc.GetUserList()
	if err != nil {
		// 处理错误，打印日志,并处理下一次请求
		log.Printf("error: %T %v\n", errors.Cause(err), errors.Cause(err))
		log.Printf("stack trace:\n %+v \n", err)
		return
	}
	//
	fmt.Println(userlist)
	return

}
