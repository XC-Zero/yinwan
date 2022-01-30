package errs

type ErrorCode int

//goland:noinspection GoSnakeCaseUsage,GoSnakeCaseUsage,GoSnakeCaseUsage,GoSnakeCaseUsage,GoSnakeCaseUsage,GoSnakeCaseUsage,GoSnakeCaseUsage
const (
	SUCCESS         = 200
	INTERNAL_ERROR  = 500
	REQUEST_ERROR   = 400
	NO_AUTH_ERROR   = 401
	FORBIDDEN_ERROR = 403
	GATE_WAY_ERROR  = 404
	RESOURCE_ERROR  = 405
	TIME_OUT_ERROR  = 408
)

func (e ErrorCode) Display() string {
	switch e {
	case SUCCESS:
		return "成功"
	case INTERNAL_ERROR:
		return "系统内部异常"
	case REQUEST_ERROR:
		return "请求参数有误"
	case NO_AUTH_ERROR:
		return "未被许可的访问"
	case FORBIDDEN_ERROR:
		return "拒绝访问"
	case GATE_WAY_ERROR:
		return "页面不存在或路由错误"
	case RESOURCE_ERROR:
		return "资源不存在"
	case TIME_OUT_ERROR:
		return "请求超时"
	default:
		return "未知错误类型"
	}
}
