package web_config_model

type Sheet struct {
	SheetName   string
	SheetColumn []SheetColumn
}

// SheetColumn 表头
type SheetColumn struct {
	SheetColumnName   string       `json:"sheet_column_name"`   // 表头名
	SheetColumnParent *SheetColumn `json:"sheet_column_parent"` // 父表头
	SheetColumnLevel  int          `json:"sheet_column_level"`  // 层级
}

// SheetCellData 每一个单元格
type SheetCellData struct {
	CellData  string  `json:"cell_data"`
	FontColor string  `json:"font_color"`
	Action    *Action `json:"action"`
}
