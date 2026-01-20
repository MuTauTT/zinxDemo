package main

import "zinxDemo/znet"

func main() {
	//1.创建Server句柄,基于Zinx的api
	s := znet.NewSever("[Zinx]V0.1")
	//2.启动server
	s.Serve()
}
