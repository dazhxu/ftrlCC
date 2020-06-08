package train

import "math"

type LR struct {
}

/**
决策函数为sigmod函数
@param w: 特征权重
@param x: 数据
@return : sigmod值
*/
func (lr *LR) Fn(w []float64, x []float64) float64 {
	var sum float64
	for idx, item := range x {
		sum += item * w[idx]
	}
	return 1.0 / (1.0 + math.Exp(-sum))
}

/**
交叉熵损失函数
*/
func (lr *LR) Loss(y float64, y_hat float64) float64 {
	return -y*math.Log(y_hat) - (1-y)*math.Log(1-y_hat)
}

/**
交叉熵损失函数对权重w的一阶导数
*/
func (lr *LR) Grad(y float64, y_hat float64, x float64) float64 {
	return (y_hat - y) / x
}
