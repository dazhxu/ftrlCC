package train

type Model interface {
	// ml模型的预测函数
	Fn(w []float64, x []float64) float64
	// ml模型的交叉熵损失函数
	Loss(y float64, y_hat float64) float64
	// ml模型的交叉熵损失函数对权重w的一阶导数
	Grad(y float64, y_hat float64, x float64) float64
}
