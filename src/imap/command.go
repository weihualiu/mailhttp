//

package imap

import (
)

// imap操作命令定义
type Command struct {
	Name string
}

// login command
type CmdLogin struct {
	Command
	UserName, Password string
}

func NewCmdLogin(username, password string) *CmdLogin {
	cmdLogin := new(CmdLogin)
	cmdLogin.Name = "SELECT"
	cmdLogin.UserName = username
	cmdLogin.Password = password
	return cmdLogin
}

func (this CmdLogin) ToByte() []byte {
	cmd := this.Name + " " + this.UserName + " " + this.Password
	return []byte(cmd)
}

type CmdList struct {
	Command
	Path string
	Regex string
}

func NewCmdList(path, regex string) *CmdList {
	cmdList := new(CmdList)
	cmdList.Name = "LIST"
	cmdList.Path = path
	cmdList.Regex = regex
	return cmdList
}

func (this CmdList) ToByte() []byte {
	cmd := this.Name + " " + this.Path + " " + this.Regex
	return []byte(cmd)
}

type CmdSelect struct {
	Command
	BoxName string
}

func NewCmdSelect(box string) *CmdSelect {
	cmdSelect := new(CmdSelect)
	cmdSelect.Name = "SELECT"
	cmdSelect.BoxName = box
	return cmdSelect
}

func (this CmdSelect) ToByte() []byte {
	cmd := this.Name + " " + this.BoxName
	return []byte(cmd)
}

type CmdFetch struct {
	Command
	NumRange string
	Summary string
}

func NewCmdFetch(numRange, summary string) *CmdFetch {
	cmdFetch := new(CmdFetch)
	cmdFetch.Name = "FETCH"
	cmdFetch.NumRange = numRange
	cmdFetch.Summary = summary
	return cmdFetch
}

func (this CmdFetch) ToByte() []byte {
	cmd := this.Name + " " + this.NumRange + " " + this.Summary
	return []byte(cmd)
}

type CmdNoop struct {
	Command
}

func NewCmdNoop() *CmdNoop {
	cmdNoop := new(CmdNoop)
	cmdNoop.Name = "NOOP"
	return cmdNoop
}

func (this CmdNoop) ToByte() []byte {
	return []byte(this.Name)
}

type CmdLogout struct {
	Command
}

func NewCmdLogout() *CmdLogout {
	cmdLogout := new(CmdLogout)
	cmdLogout.Name = "LOGOUT"
	return cmdLogout
}

func (this CmdLogout) ToByte() []byte {
	return []byte(this.Name)
}

type CmdStore struct {
	Command
	MailId string
	Flags string
	Attributes string
}

func NewCmdStore(mailid, flags, attr string) *CmdStore {
	cmdStore := new(CmdStore)
	cmdStore.Name = "STORE"
	cmdStore.MailId = mailid
	cmdStore.Flags = flags
	cmdStore.Attributes = attr
	return cmdStore
}

func (this CmdStore) ToByte() []byte {
	cmd := this.Name + " " + this.MailId + " " + this.Flags + " " + this.Attributes
	return []byte(cmd)
}

type CmdClose struct {
	Command
}

func NewCmdClose() *CmdClose {
	cmdClose := new(CmdClose)
	cmdClose.Name = "CLOSE"
	return cmdClose
}

func (this CmdClose) ToByte() []byte {
	return []byte(this.Name)
}

//EXAMINE
//SEARCH
//APPEND
//EXPUNGE
