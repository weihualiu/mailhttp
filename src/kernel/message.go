
package kernel

type MsgType int

const (
	IMAP = iota // default 0(int)
	SMTP
	WMSRV
	WS //webservice
)

// 负责与worker交互的struct
// 根据userCode和msgType取出要操作的具体实例对象
type Worker struct {
	UserInfo *UserInfo // user all information
	Type MsgType
	// 通过类型断言转换类型
	Instance interface{} // struct Imap|Smtp|Wmsrv
}

// 用户基本信息
type UserInfo struct {
	UserCode string
	UserName string
	LogonName string
	PasswdEncrypt string
}

func NewUserInfo(userCode, userName, logonName, passwdEncrypt string) *UserInfo {
	userInfo := new(UserInfo)
	userInfo.UserCode = userCode
	userInfo.UserName = userName
	userInfo.LogonName = logonName
	userInfo.PasswdEncrypt = passwdEncrypt
	return userInfo
}

func (this *UserInfo) UpdatePasswd(passwd string) {
	this.PasswdEncrypt = passwd
}

