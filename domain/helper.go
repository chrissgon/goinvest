package domain

import "math"

func ValuePerShare(value float64, shares int) float64 {
	return math.Floor((value/float64(shares))*100) / 100
}
