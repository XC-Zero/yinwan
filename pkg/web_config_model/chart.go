package web_config_model

type ChartType ComponentType

const (
	LINE_CHART ChartType = ""
	PIPE_CHART ChartType = ""
	BAR_CHART  ChartType = ""
)

type LineChart struct {
}

func (l LineChart) ChartName() {

}
