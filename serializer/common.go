package serializer

import (
	"gin_mall/pkg/e"
)

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// DataList 带有总数的Data结构
type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

// TokenData 带有token的Data结构
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

// BuildList 带有总数的列表序列化器
func BuildList(data interface{}, total uint) Response {
	return Response{
		Status: e.SUCCESS,
		Data: DataList{
			Item:  data,
			Total: total,
		},
		Msg: "ok",
	}
}
