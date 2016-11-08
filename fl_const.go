package aliim

const (
	OpenImUserAdd    string = "taobao.openim.users.add"
	OpenImUserUpdate string = "taobao.openim.users.update"
	OpenImUserDelete string = "taobao.openim.users.delete"
)

var AppKey string
var AppSecret string

// 获取服务端http地址， isPro true:正式环境 false: 沙箱环境
func GetHttpServerAddr(isPro bool) string {
	if isPro {
		return "http://gw.api.taobao.com/router/rest"
	} else {
		return "http://gw.api.tbsandbox.com/router/rest"
	}
}

// 获取服务端https地址， isPro true:正式环境 false: 沙箱环境
func GetHttpsServerAddr(isPro bool) string {
	if isPro {
		return "https://eco.taobao.com/router/rest"
	} else {
		return "https://gw.api.tbsandbox.com/router/rest"
	}
}
