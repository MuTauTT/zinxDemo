package znet

import (
	"fmt"

	"zinxDemo/ziface"
)

// iServer的接口实现,定义一个Server的服务器模块
type Server struct {
	//服务器的名称
	Name string
	//服务器绑定的ip类型
	IPVersion string
	//服务器的IP地址
	IP string
	//服务器监听的端口
	Port int
}

// 启动服务器
func (s *Server) Start() {
	fmt.Println()
}

// 停止服务器
func (s *Server) Stop() {

}

// 运行服务器
func (s *Server) Serve() {

}

/*
初始化Server模块的方法
*/
func NewSever(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
