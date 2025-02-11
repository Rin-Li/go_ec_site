package v1

import (
	"encoding/json"
	"gin-mall-tmp/serializer"
)

func ErrorResPonse(err error) serializer.Response{
	if _,ok := err.(*json.UnmarshalTypeError);ok{
		return serializer.Response{
			Status: 400,
			Msg: "JSON not comapre",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Status: 400,
		Msg: "Paremeter wrong",
		Error: err.Error(),
	}
}