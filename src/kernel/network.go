package kernel

import (
	"net"
	"time"
)

type Network interface {
	Connect() (*net.TCPConn, error)
	ReConnect() (*net.TCPConn, error)
	Closed() error
	SendAndReceive([]byte) ([]byte, error) 
	SetEnv() error
	//异常处理
	HandleError()
	Command(message string) string
}

type Net struct {
	//连接句柄
	Sock *net.TCPConn
	//时间戳，最后一次交互
	Snapshot time.Time
	//是否SSL
	Ssl bool
	//超时时间
	Timeout int64
	//服务端地址串 localhost:9827
	Addr string
	// 交互通道
	PutCmd chan string
	GetCmd chan string
	CloseCmd chan int
}
