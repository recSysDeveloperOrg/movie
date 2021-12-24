package constant

import "fmt"

type ErrorCode struct {
	Code int64
	Msg  string
}

func makeErrorCode(code int64, msg string) *ErrorCode {
	return &ErrorCode{
		Code: code,
		Msg:  msg,
	}
}

func BuildErrCode(detail interface{}, errCode *ErrorCode) *ErrorCode {
	return makeErrorCode(errCode.Code, fmt.Sprintf(errCode.Msg, detail))
}

var (
	RetSuccess            = makeErrorCode(0, "成功")
	RetParamsErr          = makeErrorCode(1, "参数非法:%v")
	RetNoSuchMovieID      = makeErrorCode(2, "没有该ID为:%v对应的电影信息")
	RetOperationForbidden = makeErrorCode(3, "不支持添加此过滤规则")
	RetWriteRepoErr       = makeErrorCode(100, "写库错误:%v")
	RetReadRepoErr        = makeErrorCode(101, "查库错误:%v")
	RetRpcCallFailed      = makeErrorCode(200, "调用下游错误")
	RetSysErr             = makeErrorCode(999, "系统未知错误:%v")
)
