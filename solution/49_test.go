package solution

import (
	"fmt"
	"testing"
)

func TestSortString(t *testing.T) {
	a := "cba"
	b := sortString(a)
	fmt.Println(a, b)
}
