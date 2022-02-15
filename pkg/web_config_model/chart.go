package web_config_model

type ChartType ComponentType

//goland:noinspection GoSnakeCaseUsage,GoSnakeCaseUsage,GoSnakeCaseUsage
const (
	LINE_CHART ChartType = iota + 40001
	PIPE_CHART
	BAR_CHART
)

type LineChart struct {
	LineChartLabel string   `json:"line_chart_label"`
	XAxisData      []string `json:"x_axis_data"`
	Legend         []string `json:"legend"`
	LineChartData  []map[string]interface{}
}

func (l LineChart) ComponentName() string {
	return "LineChart"
}

func (l LineChart) ComponentType() ComponentType {
	return ComponentType(LINE_CHART)
}

func (l LineChart) ChildComponentFunc() []WebComponent {
	return nil
}

func (l LineChart) ChartName() string {
	return l.LineChartLabel
}

// ChartXAxisData X轴坐标数据
func (l LineChart) ChartXAxisData() map[string]interface{} {
	return map[string]interface{}{
		"type": "category",
		"data": l.XAxisData,
		"axisTick": map[string]interface{}{
			"alignWithLabel": true,
		},
	}
}

// ChartYAxisData Y轴坐标数据
func (l LineChart) ChartYAxisData() map[string]interface{} {
	return nil
}

// ChartData 图表数据
func (l LineChart) ChartData() map[string]interface{} {
	return map[string]interface{}{
		"series": l.LineChartData,
	}
}
