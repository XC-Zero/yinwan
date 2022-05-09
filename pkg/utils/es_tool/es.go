package es_tool

import (
	"fmt"
	"github.com/olivere/elastic/v7"
)

func ESDocToUpdateScript(doc map[string]interface{}) *elastic.Script {

	baseStr := "ctx._source.%s=params.%s;"
	scriptsUrl := ""
	for key := range doc {
		scriptsUrl += fmt.Sprintf(baseStr, key, key)
	}
	return elastic.NewScriptInline(scriptsUrl).
		Params(doc)
}
