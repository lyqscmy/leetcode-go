package solution

import (
	"fmt"
	"testing"
)

func TestSubsets(t *testing.T) {
	res := subsets([]int{1, 2, 3})
	fmt.Println(res)
}
