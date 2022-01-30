package _interface

type Staff interface {
	// Login 登录
	Login() (tokenPtr *string, errorMessage string)
	// LogOut 退出登录
	LogOut()
}
