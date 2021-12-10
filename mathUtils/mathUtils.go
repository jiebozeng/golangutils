package mathUtils

import (
	"math"
)

//取float小数点后N位
func GetFloatNBit(source float64,n int32) float64 {
	nten := math.Pow(10,float64(n))
	return math.Round(source*nten)/nten
}
