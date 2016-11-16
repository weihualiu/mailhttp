
package services

import (
	"fmt"
	"sync"
)

import "kernel"
import "imap"

// 管理用户信息及关联的Imap、Smtp、Http等协议
// 提供查询对应实例、增加实例、删除实例的接口

type workerMapT map[string]map[kernel.MsgType]*kernel.Worker

var workerMap workerMapT
var workerMtx sync.RWMutex

// 包引入的时候启动变量初始化
func init() {
	fmt.Println("router init")
	workerMap = make(workerMapT, 1000000)
}

func getInst(userCode string, typeMsg kernel.MsgType) interface{} {
	workerMtx.RLock()
	defer workerMtx.RUnlock()
	workers, ok := workerMap[userCode]
	if !ok {
		return nil
	}
	worker, ok := workers[typeMsg]
	if !ok {
		return nil
	}
	return worker.Instance
}

func GetImap(userCode string) *imap.ImapNet {
	inf := getInst(userCode, kernel.IMAP)
	if inf == nil {
		//fmt.Println("get imap failed!")
		return nil
	}
	imapNet, ok := inf.(*imap.ImapNet)
	if !ok {
		//fmt.Println("get imap convert failed!")
		return nil
	}
	return imapNet
}

func SetInst(worker *kernel.Worker) error {
	userCode := worker.UserInfo.UserCode
	msgType := worker.Type
	
	workerMtx.Lock()
	defer workerMtx.Unlock()
	
	workerm, ok := workerMap[userCode]
	if !ok {
		workerm = make(map[kernel.MsgType]*kernel.Worker)
	}
	workerm[msgType] = worker
	workerMap[userCode] = workerm
	
	return nil
}

func DelUser(userCode string) error {
	workerMtx.Lock()
	defer workerMtx.Unlock()
	// 清除之前，清理相关goroutine、channel
	delete(workerMap, userCode)
	return nil
}

func DelImap(userCode string) error {
	workerMtx.Lock()
	defer workerMtx.Unlock()
	
	workers := workerMap[userCode]
	delete(workers, kernel.IMAP)
	if len(workers) == kernel.IMAP {
		delete(workerMap, userCode)
	}else{
		workerMap[userCode] = workers
	}
	return nil
}

