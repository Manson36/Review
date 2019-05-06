package message

const (
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"

	)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

//定义两个消息。。后面需要再添加
type LoginMes struct {
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code int `json:"code"`
	Error string `json:"error"`
}