package _const

type GinErrorCode int

const (
	CONTINUE                          GinErrorCode = 100
	OK                                             = 200
	UNAUTHORIZED_ERROR                             = 401
	FORBIDDEN_ERROR                                = 403
	PAGE_NOT_FOUND_ERROR                           = 404
	MethodNotAllowedError                          = 405
	TIME_OUT_ERROR                                 = 408
	CONFLICT_ERROR                                 = 409
	REQUEST_PARM_ERROR                             = 411
	REQUEST_RANGE_NOT_SATISFIED_ERROR              = 416
	EXPECTATION_FAILED_ERROR                       = 417
	UPGRADE_REQUIRED_ERROR                         = 426
	TOO_MANY_REQUEST_ERROR                         = 429
	INTERNAL_ERROR                                 = 500
	NETWORK_CONNECT_TIMEOUT_ERROR                  = 599
)

func (ge GinErrorCode) Display() string {
	switch ge {
	case CONTINUE:
		return "继续"
	case OK:
		return "成功"
	case UNAUTHORIZED_ERROR:
		return "抱歉，您没有权限访问呢"
	case FORBIDDEN_ERROR:
		return "您被拒绝访问了"
	case PAGE_NOT_FOUND_ERROR:
		return "哎呀，页面走丢了"
	case MethodNotAllowedError:
		return "您的请求方式有误哦"
	case TIME_OUT_ERROR:
		return "访问超时，请查看服务器网络呢"
	case CONFLICT_ERROR:
		return "出现冲突了呢"
	case REQUEST_PARM_ERROR:
		return "请求参数有误哦"
	case REQUEST_RANGE_NOT_SATISFIED_ERROR:
		return "请求范围没办法被满足诶"
	case EXPECTATION_FAILED_ERROR:
		return "这个错误果然在预料之内"
	case UPGRADE_REQUIRED_ERROR:
		return "系统版本要升级了"
	case TOO_MANY_REQUEST_ERROR:
		return "请求量太大了，扩容下服务器吧"
	case INTERNAL_ERROR:
		return "系统内部错误"
	case NETWORK_CONNECT_TIMEOUT_ERROR:
		return "网络连接超时"
	default:
		return "其他类型错误，麻烦联系管理员哦"
	}
}
