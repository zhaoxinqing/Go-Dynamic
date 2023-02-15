package lib

import "math"

// AddFloat64 float64 相加
func AddFloat64(n1 float64, n2 float64, retain int) (result float64) {
	return Unwrap(Wrap(n1, retain)+Wrap(n2, retain), retain)
}

//将f loat64 转成精确的 int64
func Wrap(num float64, retain int) int64 {
	return int64(num * math.Pow10(retain))
}

//将 int64 恢复成正常的 float64
func Unwrap(num int64, retain int) float64 {
	return float64(num) / math.Pow10(retain)
}

//精准 float64
func WrapToFloat64(num float64, retain int) float64 {
	return num * math.Pow10(retain)
}

//精准 int64
func UnwrapToInt64(num int64, retain int) int64 {
	return int64(Unwrap(num, retain))
}
