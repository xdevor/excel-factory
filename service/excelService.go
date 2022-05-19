package service

import (
	"echoException/excel-faker/helper"
	"fmt"
	"strconv"

	excelize "github.com/xuri/excelize/v2"
)

func SaveExcel(data [][]string, filename string) (ok bool) {
	f := excelize.NewFile()
	sheet := "Sheet1"
	index := f.NewSheet(sheet)
	rowLen := len(data)

	for i := 0; i < rowLen; i++ {
		for j := 0; j < len(data[i]); j++ {
			axis := helper.ConvertIntToExcelTitle(j+1) + strconv.Itoa(i+1)
			f.SetCellValue(sheet, axis, data[i][j])
		}
	}

	f.SetActiveSheet(index)
	if err := f.SaveAs(filename); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
