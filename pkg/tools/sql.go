package tools

var contrastTable map[string]string

func init() {
	contrastTable = make(map[string]string, 0)
	contrastTable["update"] = "更新"
	contrastTable["create"] = "创建"
	contrastTable["delete"] = "删除"
	contrastTable["table"] = "表"
	contrastTable["schema"] = "账套"

}
