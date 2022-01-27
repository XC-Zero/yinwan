package web_config_model

type ChildComponentFunc func() ([]WebComponent, error)

// WebComponent web组件
type WebComponent interface {
	ComponentName() string
	ChildComponentFunc() ChildComponentFunc
}

// Chart 图表
type Chart interface {
	WebComponent
	ChartType() ChartType
	ChartName() string
}

//
type Block interface {
	WebComponent
	BlockType() BlockType
	BlockName() string
}

// WebRootComponent 页面配置/根组件
type WebRootComponent struct {
	WebComponentName string
	WebViewName      string
	WebContent       []WebComponent
}

func (wrc WebRootComponent) ComponentName() string {
	return wrc.WebComponentName
}

func (wrc WebRootComponent) ChildComponentFunc() ChildComponentFunc {
	return func() ([]WebComponent, error) {
		return wrc.WebContent, nil
	}
}
