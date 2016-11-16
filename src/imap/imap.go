
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
	
}

func NewImapNet(ssl bool, timeout int64, addr string) *ImapNet {
	imapNet := new(ImapNet)
	imapNet.Ssl = ssl
	imapNet.Timeout = timeout
	imapNet.Addr = addr
	imapNet.PutCmd = make(chan string)
	imapNet.GetCmd = make(chan string)
	imapNet.CloseCmd = make(chan int)
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
	close(this.PutCmd)
	close(this.GetCmd)
	this.CloseCmd <- 0
	close(this.CloseCmd)
	
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
	this.PutCmd <- message
	return <-this.GetCmd
}

func (this *ImapNet) Instance() error {
	go func(in *ImapNet) {
		for{
			select {
				case <-in.PutCmd: // manage layer to net
					// email data
					tmp := "0"
					in.GetCmd<- tmp
				case <-in.CloseCmd: // close goroutine
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

type Command struct {
	
}