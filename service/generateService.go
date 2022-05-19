package service

import (
	"echoException/excel-faker/dto"
	"math/rand"
	"time"
)

func GenerateFakeExcelData(request dto.ExportFakeExcel) [][]string {
	header := ParseHeaderByExportFakeExcelColumns(request.Columns)
	dataTypes := ParseDataTypesByExportFakeExcelColumns(request.Columns)
	output := [][]string{header}

	for i := 0; i < request.RowCount; i++ {
		var curData []string
		for j := 0; j < len(dataTypes); j++ {
			data := GetFakeData(dataTypes[j])
			curData = append(curData, data)
		}
		output = append(output, curData)
	}

	return output
}

func GenerateExcelFileName(oriFileName string) string {
	randPostfix := GenerateRandomString(6)
	return "storage/public/excel/" + oriFileName + "-" + time.Now().Format("2006-0201-15-04-05") + "-" + randPostfix + ".xlsx"
}

func GenerateRandomString(n int) string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
	str := make([]rune, n)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}
