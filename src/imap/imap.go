
package imap

import (
	"net"
	"time"
	"runtime"
	//"fmt"
)

import "kernel"

type ImapNet struct {
	kernel.Net
	// authenticate true false
	AuthFlag bool
	oldCmd *Command
	currCmd *Command
	// 操作命令序号
	seqCmd int
}

func NewImapNet(ssl bool, timeout int64, addr string) *ImapNet {
	imapNet := new(ImapNet)
	imapNet.Ssl = ssl
	imapNet.Timeout = timeout
	imapNet.Addr = addr
	imapNet.PutCh = make(chan string)
	imapNet.GetCh = make(chan string)
	imapNet.CloseCh = make(chan int)
	imapNet.AuthFlag = false
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
	// close all channel
	close(this.PutCh)
	close(this.GetCh)
	this.CloseCh <- 0
	close(this.CloseCh)
	
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
	this.PutCh <- message
	return <-this.GetCh
}

func (this *ImapNet) Instance() error {
	go func(in *ImapNet) {
		for{
			select {
				case <-in.PutCh: // manage layer to net
					// email data
					tmp := "0"
					in.GetCh <- tmp
				case <-in.CloseCh: // close goroutine
					// goroutine exit
					runtime.Goexit()
				default:
					// nothing
					//fmt.Println("imap goroutine nothing!")
			}
		}
	}(this)
	
	return nil
}

type Parse struct {
}

