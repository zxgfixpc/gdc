package main

import (
	"fmt"
	"gdc/httpserver"
)

// 项目启动入口
func main() {
	//
	fmt.Println("starting...")
	httpserver.Start()
}
