package web_config_model

type ComponentType int

//goland:noinspection GoSnakeCaseUsage
const (
	ROOT_COMPONENT  ComponentType = -1
	FIRST_COMPONENT ComponentType = iota + 1
	SECOND_COMPONENT
	THIRTY_COMPONENT
	FOURTH_COMPONENT
	FIFTH_COMPONENT
)

func (c ComponentType) Display() string {
	switch c {
	case ROOT_COMPONENT:
		return "根组件"
	case FIRST_COMPONENT:
		return "一级组件"
	case SECOND_COMPONENT:
		return "二级组件"
	case THIRTY_COMPONENT:
		return "三级组件"
	case FOURTH_COMPONENT:
		return "四级组件"
	case FIFTH_COMPONENT:
		return "五级组件"
	default:
		return "其他"
	}
}

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
