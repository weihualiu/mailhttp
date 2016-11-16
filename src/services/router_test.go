package services

import "testing"
import "kernel"
import "imap"

func TestSetInst(t *testing.T) {
	//workerMap = make(workerMapT, 1000000)
	
	worker := new(kernel.Worker)
	userInfo := new(kernel.UserInfo)
	userInfo.UserCode = "xiaobb"
	worker.UserInfo = userInfo
	worker.Type = 0
	imap := new(imap.ImapNet)
	imap.Addr = "0.0.0.0"
	imap.Ssl = true
	worker.Instance = imap
	
	if SetInst(worker) != nil {
		t.Errorf(`TestSetInst() is failed!`)
	}
}

func TestGetImap(t *testing.T) {
	imap1 := GetImap("xiaobb")
	if imap1 == nil {
		t.Errorf(`TestGetImap1() is failed!`)
	}
	
	if !imap1.Ssl {
		t.Errorf(`TestGetImap2() is failed!`)
	}
}

func TestDelImap(t *testing.T) {
	err := DelImap("xiaobb")
	//err = DelImap("xx")
	if err != nil {
		t.Errorf(`TestDelImap() is failed!`)
	}
}

func TestDelUser(t *testing.T) {
	err := DelUser("xiaobb")
	err = DelUser("xxx")
	
	if err != nil {
		t.Errorf(`TestDelUser() is failed!`)
	}
}