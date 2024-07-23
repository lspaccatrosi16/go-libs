package regression

import (
	"math"

	"github.com/lspaccatrosi16/go-libs/statistics/data"
)

type Regression struct {
	M  float64
	C  float64
	R  float64
	R2 float64
}

func (r *Regression) Y(x float64) float64 {
	return r.M*x + r.C
}

func (r *Regression) X(y float64) float64 {
	return (y - r.C) / r.M
}

func YonX(bivar data.BivariateData) Regression {
	m := bivar.Sxy / bivar.Sxx
	c := bivar.MeanY - m*bivar.MeanX

	r := bivar.Sxy / math.Sqrt(bivar.Sxx*bivar.Syy)

	return Regression{
		M:  m,
		C:  c,
		R:  r,
		R2: r * r,
	}
}

func XonY(bivar data.BivariateData) Regression {
	m := bivar.Sxy / bivar.Syy
	c := bivar.MeanY - bivar.MeanX/m

	r := bivar.Sxy / math.Sqrt(bivar.Sxx*bivar.Syy)

	return Regression{
		M:  1 / m,
		C:  c,
		R:  r,
		R2: r * r,
	}
}

func LinearRegression(m, c float64) Regression {
	return Regression{
		M:  m,
		C:  c,
		R:  0,
		R2: 0,
	}
}
