package web_config_model

type Action struct {
	ActionName    string `json:"action_name"`
	ActionType    string `json:"action_type"`
	ActionContent string `json:"action_content"`
}

// RouterAction 路由响应
type RouterAction struct {
	//action
}

// RedirectAction 跳转响应
type RedirectAction struct {
}

// ViewAction 页面响应
type ViewAction struct {
}
