package data

import "math"

type UnivariateData struct {
	N     float64
	X     []float64
	SumX  float64
	MeanX float64
	SumX2 float64
	Sxx   float64
	VarX  float64
	StdX  float64
}

func GetUnivariate(xd []float64) UnivariateData {
	n := float64(len(xd))

	var sx, sx2 float64
	for i := 0; i < len(xd); i++ {
		sx += xd[i]
		sx2 += xd[i] * xd[i]
	}

	sxx := sx2 - ((sx * sx) / n)
	vx := sxx / (n - 1)

	return UnivariateData{
		N:     n,
		X:     xd,
		SumX:  sx,
		MeanX: sx / n,
		SumX2: sx2,
		Sxx:   sxx,
		VarX:  vx,
		StdX:  math.Sqrt(vx),
	}
}
