package solution

const (
	intSize = 32 << (^uint(0) >> 63) // 32 or 64

	MinInt = -1 << (intSize - 1)
	MaxInt = 1<<(intSize-1) - 1
)

var M int
var N int

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
