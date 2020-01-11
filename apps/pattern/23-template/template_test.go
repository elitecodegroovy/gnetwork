package template

import (
	"fmt"
	"sort"
	"testing"
)

func TestTemplate(t *testing.T) {
	var myList MyList = []int{12, 7, 9, 3, 89, 3, 0, 89, 1}
	sort.Sort(myList)
	//var sortedResult = []int{0, 1, 3, 3, 7, 9, 12, 89, 89}
	c := fmt.Sprintf("%v", myList)
	if c != "[0 1 3 3 7 9 12 89 89]" {
		t.Fatal("The int array was sorted, but it failed.")
	}
}
