package dto

type ExportFakeExcel struct {
	FileName string                  `json:"fileName" form:"fileName" binding:"required"`
	RowCount int                     `json:"rowCount" form:"rowCount" binding:"required"`
	Columns  []ExportFakeExcelColumn `json:"columns" form:"columns" binding:"required"`
}

type ExportFakeExcelColumn struct {
	Name         string `json:"name" form:"name"`
	Source       string `json:"source" form:"source"`     // random or template
	DataType     string `json:"dataType" form:"dataType"` // fake data type
	Template     string `json:"template" form:"template"`
	AutoIncrease bool   `json:"autoIncrease" form:"autoIncrease"`
}
