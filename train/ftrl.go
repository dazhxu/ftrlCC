package train

import "math"

type FTRL struct {
	Dim          int
	DecisionFunc *LR
	Z            []float64
	N            []float64
	W            []float64
	L1           float64
	L2           float64
	Alpha        float64
	Beta         float64
}

func Init(dim int, l1 float64, l2 float64, alpha float64, beta float64, decisionFunc *LR) *FTRL {
	ftrl := new(FTRL)
	ftrl.Dim = dim
	ftrl.DecisionFunc = decisionFunc
	ftrl.L1 = l1
	ftrl.L2 = l2
	ftrl.Alpha = alpha
	ftrl.Beta = beta
	ftrl.Z = make([]float64, dim)
	ftrl.N = make([]float64, dim)
	ftrl.W = make([]float64, dim)
	return ftrl
}

/**
预测
*/
func (ftrl *FTRL) Predict(x []float64) float64 {
	return ftrl.DecisionFunc.Fn(ftrl.W, x)
}

/**
更新
*/
func (ftrl *FTRL) Update(x []float64, y float64) float64 {
	for i := 0; i < ftrl.Dim; i++ {
		if math.Abs(ftrl.Z[i]) <= ftrl.L1 {
			ftrl.W[i] = 0
		} else {
			ftrl.W[i] = (Sign(ftrl.Z[i])*ftrl.L1 - ftrl.Z[i]) / (ftrl.L2 + (ftrl.Beta+math.Sqrt(ftrl.N[i]))/ftrl.Alpha)
		}
	}

	y_hat := ftrl.Predict(x)

	for i := 0; i < ftrl.Dim; i++ {
		g := ftrl.DecisionFunc.Grad(y, y_hat, x[i])
		sigma := (math.Sqrt(ftrl.N[i]+g*g) - math.Sqrt(ftrl.N[i])) / ftrl.Alpha
		ftrl.Z[i] += g - sigma*ftrl.W[i]
		ftrl.N[i] += g * g
	}

	return ftrl.DecisionFunc.Loss(y, y_hat)

}
