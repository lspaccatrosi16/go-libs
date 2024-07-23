package data

import "github.com/lspaccatrosi16/go-libs/internal/pkgError"

type BivariateData struct {
	N     float64
	X     []float64
	Y     []float64
	SumX  float64
	SumY  float64
	MeanX float64
	MeanY float64
	SumX2 float64
	SumY2 float64
	SumXY float64
	Sxx   float64
	Syy   float64
	Sxy   float64
}

var errorf = pkgError.ErrorfFactory("statistics/data")

func GetBivariate(xd, yd []float64) (BivariateData, error) {
	if len(xd) != len(yd) {
		return BivariateData{}, errorf("length of x and y data must be equal")
	}

	var sx, sy, sx2, sy2, sxy float64
	n := float64(len(xd))

	for i := 0; i < len(xd); i++ {
		sx += xd[i]
		sy += yd[i]
		sx2 += xd[i] * xd[i]
		sy2 += yd[i] * yd[i]
		sxy += xd[i] * yd[i]
	}

	return BivariateData{
		N:     n,
		X:     xd,
		Y:     yd,
		SumX:  sx,
		SumY:  sy,
		MeanX: sx / n,
		MeanY: sy / n,
		SumX2: sx2,
		SumY2: sy2,
		SumXY: sxy,
		Sxx:   sx2 - ((sx * sx) / n),
		Syy:   sy2 - ((sy * sy) / n),
		Sxy:   sxy - ((sx * sy) / n),
	}, nil

}
