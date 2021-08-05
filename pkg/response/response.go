package response

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/lidongyooo/swag-blog-api/pkg/strings"
)

type Response struct {
	Code int         `json:"code"` // 错误码
	Msg  string      `json:"message"`  // 错误描述
	Data interface{} `json:"data"` // 返回数据
}

type Errors struct {
	Code int         `json:"code"` // 错误码
	Msg  string      `json:"message"`  // 错误描述
	Errors map[string]string `json:"errors"`
}

func (res *Response) WithMsg(message string) Response {
	return Response{
		Code: res.Code,
		Msg:  message,
		Data: res.Data,
	}
}

func (res *Response) WithData(data interface{}) Response {
	return Response{
		Code: res.Code,
		Msg:  res.Msg,
		Data: data,
	}
}

// ToString 返回 JSON 格式的字符串
func (res *Response) ToString() string {
	raw, _ := json.Marshal(res)
	return string(raw)
}

func Error(code int, err error) *Errors {
	var (
		msg string
		errors = make(map[string]string)
	)

	if _, ok := err.(validator.ValidationErrors); ok {
		msg = "The given data was invalid."

		for _, err := range err.(validator.ValidationErrors) {
			errors[strings.SnakeString(err.Field())] = err.Tag()
		}
	} else {
		msg = err.Error()
	}

	return &Errors{
		Code: code,
		Msg: msg,
		Errors: errors,
	}
}

func New(code int, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}
