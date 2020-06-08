package train

/**
模拟numpy的sign函数
*/
func Sign(num float64) float64 {
	if num > 0 {
		return 1
	} else if num < 0 {
		return -1
	} else {
		return 0
	}
}
