
package imap

import (
	"net"
	"time"
)

import "kernel"

type ImapNet struct {
	kernel.Net
	// authenticate true false
	AuthFlag bool
	// 交互通道
	Command chan string
}

func NewImapNet(ssl bool, timeout int64, addr string) *ImapNet {
	imapNet := new(ImapNet)
	imapNet.Ssl = ssl
	imapNet.Timeout = timeout
	imapNet.Addr = addr
	imapNet.Command = make(chan string)
	return imapNet
}

func (this *ImapNet) Connect() (error) {
	//DaiTcp
	tcpAddr , err := net.ResolveTCPAddr("tcp4", this.Addr)
	if err != nil {
		return err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	}
	
	this.Sock = conn
	this.Snapshot = time.Now()
	
	return nil
}

func (this *ImapNet) ReConnect() (error) {
	return nil
}

func (this *ImapNet) Closed() error {
	this.Sock.Close()
	return nil
}

// 数据收发
func (this *ImapNet) SendAndReceive() ([]byte, error) {
	return nil, nil
}

func (this *ImapNet) SetEnv() error {
	return nil
}

func (this *ImapNet) HandleError() {
}

func (this *ImapNet) Parse() {
}

func (this *ImapNet) Build() {
}

func (this *ImapNet) Process(message string) string {
	this.Command <- message
	return <-this.Command
}

func (this *ImapNet) Instance() {
	go func() {
		
	}()
}

type Command struct {
	
}