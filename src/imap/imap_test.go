
package imap

import "testing"

import "kernel"
import "fmt"
import "time"

func TestConnect(t *testing.T) {
	worker := new(kernel.Worker)
	imapc := NewImapNet(false, 600, "182.119.175.196:143")
	worker.Instance = imapc
	value, ok := worker.Instance.(*ImapNet)
	if ok {
		fmt.Println("addr: ", value.Addr)
	}else{
		fmt.Println("error")
	}
	
	err := imapc.Connect()
	if err != nil {
		t.Errorf(`TestConnect() is failed!`, err)
	}
}

func TestInstance(t *testing.T) {
	imapc := NewImapNet(false, 600, "182.119.175.196:143")
	err := imapc.Instance()
	time.Sleep(time.Second)
	imapc.CloseCh <- 0 // goroutine close
	fmt.Println("imap goroutine exit!")
	time.Sleep(time.Second)
	if err != nil {
		t.Errorf(`TestInstance() is failed!`, err)
	}
	
}
