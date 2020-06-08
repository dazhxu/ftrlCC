package models

// 数据模型
type DataEntry struct {
	X []float64 `json:"x"`
	Y float64   `json:"y"`
}

// 数据记录
type Record struct {
	MspId string
	Count int
}

// 返回信息
type Response struct {
	Retcode int    `json:"retcode"`
	Retmsg  string `json:"retmsg"`
}
