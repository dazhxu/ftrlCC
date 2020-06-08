package train

import "math"

type FTLR struct {
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

func Init(dim int, l1 float64, l2 float64, alpha float64, beta float64, decisionFunc *LR) *FTLR {
	ftlr := new(FTLR)
	ftlr.Dim = dim
	ftlr.DecisionFunc = decisionFunc
	ftlr.L1 = l1
	ftlr.L2 = l2
	ftlr.Alpha = alpha
	ftlr.Beta = beta
	ftlr.Z = make([]float64, dim)
	ftlr.N = make([]float64, dim)
	ftlr.W = make([]float64, dim)
	return ftlr
}

/**
预测
*/
func (ftlr *FTLR) Predict(x []float64) float64 {
	return ftlr.DecisionFunc.Fn(ftlr.W, x)
}

/**
更新
*/
func (ftlr *FTLR) Update(x []float64, y float64) float64 {
	for i := 0; i < ftlr.Dim; i++ {
		if math.Abs(ftlr.Z[i]) <= ftlr.L1 {
			ftlr.W[i] = 0
		} else {
			ftlr.W[i] = (Sign(ftlr.Z[i])*ftlr.L1 - ftlr.Z[i]) / (ftlr.L2 + (ftlr.Beta+math.Sqrt(ftlr.N[i]))/ftlr.Alpha)
		}
	}

	y_hat := ftlr.Predict(x)

	for i := 0; i < ftlr.Dim; i++ {
		g := ftlr.DecisionFunc.Grad(y, y_hat, x[i])
		sigma := (math.Sqrt(ftlr.N[i]+g*g) - math.Sqrt(ftlr.N[i])) / ftlr.Alpha
		ftlr.Z[i] += g - sigma*ftlr.W[i]
		ftlr.N[i] += g * g
	}

	return ftlr.DecisionFunc.Loss(y, y_hat)

}
