package solution

import (
	"fmt"
	"testing"
)

func Test_reconstructQueue(t *testing.T) {
	people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	people = reconstructQueue(people)
	fmt.Println(people)
}
