package response


var MsgFlags = map[int]string {
	SUCCESS:                 "ok",
	ERROR:                   "fail",
	INVALID_PARAMS:          "请求参数错误",
	ERROR_NOT_EXIST_ARTICLE: "该文章不存在",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
