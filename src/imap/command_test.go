package imap

import "testing"
import "fmt"

func TestToByte(t *testing.T) {
	cmdLogin := NewCmdLogin("xiaobb", "password")
	data := cmdLogin.ToByte()
	if data == nil {
		t.Errorf(`cmd login is failed!`)
	}
	fmt.Println("cmd login:",string(data))
	
	cmdList := NewCmdList("\"\"", "*")
	data = cmdList.ToByte()
	if data == nil {
		t.Errorf(`cmd list is failed!`)
	}
	fmt.Println("cmd list:", string(data))
	
	cmdSelect := NewCmdSelect("inbox")
	data = cmdSelect.ToByte()
	if data == nil {
		t.Errorf(`cmd select is failed!`)
	}
	fmt.Println("cmd select:", string(data))
	
	cmdFetch := NewCmdFetch("1:10", "body[text]")
	data = cmdFetch.ToByte()
	if data == nil {
		t.Errorf(`cmd fetch is failed!`)
	}
	fmt.Println("cmd fetch:", string(data))
	
	cmdNoop := NewCmdNoop()
	data = cmdNoop.ToByte()
	if data == nil {
		t.Errorf(`cmd noop is failed!`)
	}
	fmt.Println("cmd noop:", string(data))
	
	cmdLogout := NewCmdLogout()
	data = cmdLogout.ToByte()
	if data == nil {
		t.Errorf(`cmd logout is failed!`)
	}
	fmt.Println("cmd logout:", string(data))
}