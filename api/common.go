package api

import (
	"encoding/json"
	"gin_mall/consts"
	"gin_mall/serializer"
)

func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: consts.IlleageRequest,
			Msg:    "JSON类型不匹配",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: consts.IlleageRequest,
		Msg:    "参数错误",
		Error:  err.Error(),
	}
}
