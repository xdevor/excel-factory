package service

import "echoException/excel-faker/dto"

func ParseHeaderByExportFakeExcelColumns(columns []dto.ExportFakeExcelColumn) []string {
	var header []string
	for i := 0; i < len(columns); i++ {
		curColumn := columns[i]
		header = append(header, curColumn.Name)
	}
	return header
}

func ParseDataTypesByExportFakeExcelColumns(columns []dto.ExportFakeExcelColumn) []string {
	var dataTypes []string
	for i := 0; i < len(columns); i++ {
		curColumn := columns[i]
		dataTypes = append(dataTypes, curColumn.DataType)
	}
	return dataTypes
}
