
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
	UserCode string
	Type MsgType
	// 通过类型断言转换类型
	Instance interface{} // struct Imap|Smtp|Wmsrv
}


