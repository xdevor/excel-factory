package helper

func ConvertIntToExcelTitle(n int) string {
	result := ""
	for {
		if n <= 26 {
			result = string(n+64) + result
			break
		} else {
			result = string((n-1)%26+1+64) + result
			n = (n - 1) / 26
		}
	}
	return result
}
