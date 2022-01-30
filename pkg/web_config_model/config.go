package web_config_model

type ComponentType string

//goland:noinspection GoSnakeCaseUsage
const ROOT_COMPONENT ComponentType = "ROOT"

// WebComponent web组件
type WebComponent interface {
	ComponentName() string
	ComponentType() ComponentType
	ChildComponentFunc() []WebComponent
}

// Chart 图表
type Chart interface {
	WebComponent
	// ChartName 图表标题
	ChartName() string
	// ChartXAxisData X轴坐标数据
	ChartXAxisData() map[string]interface{}
	// ChartYAxisData Y轴坐标数据
	ChartYAxisData() map[string]interface{}
	// ChartData 图表数据
	ChartData() map[string]interface{}
}

// Block 区块
type Block interface {
	WebComponent
	BlockName() string
}

// StepList 步骤/流程
type StepList interface {
	WebComponent
	StepListNumber() int
}

// WebRootComponent 页面配置/根组件
// 每一个页面都有一个这玩意
// todo 要不要把所有的该结构体的对象管理起来？ 内存map？ 数据库表？
type WebRootComponent struct {
	WebComponentName string
	WebViewName      string
	WebContent       []WebComponent
}

func (wrc WebRootComponent) ComponentName() string {
	return wrc.WebComponentName
}

func (wrc WebRootComponent) ChildComponentFunc() []WebComponent {
	return wrc.WebContent
}

func (wrc WebRootComponent) ComponentType() ComponentType {
	return ROOT_COMPONENT
}
