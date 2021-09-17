package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

type Server struct {
	// 必需配置，无默认值
	IP   string
	Port int
	// 必需配置，有默认值
	Protocol string
	Timeout  time.Duration
	Maxconns int
	// 可选配置
	TLS *tls.Config
}

// 定义一个函数类型，接收*Server为参数
type Option func(*Server)

// 设置协议的高阶函数方法
func Protocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}

func Timeout(t time.Duration) Option {
	return func(s *Server) {
		s.Timeout = t
	}
}

func Maxconns(mc int) Option {
	return func(s *Server) {
		s.Maxconns = mc
	}
}

func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

func NewServer(ip string, port int, options ...Option) (*Server, error) {
	// 必需参数直接赋值/设置默认值，非必需参数可默认为类型零值
	server := &Server{
		IP:       ip,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		Maxconns: 10,
	}
	// 通过forloop遍历设置传入的参数
	for _, option := range options {
		option(server)
	}
	return server, nil
}

func main() {
	serverA, _ := NewServer("127.0.0.1", 2001)
	serverB, _ := NewServer("127.0.0.1", 2001, Protocol("udp"))
	serverC, _ := NewServer("127.0.0.1", 2001, Protocol("udp"), TLS(&tls.Config{}))
	fmt.Println(serverA, serverB, serverC)
}
