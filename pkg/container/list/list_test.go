package list

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {

	l := NewList()

	l.GenerateListInt32([]int32{23, 52, 1, 2, 43, 25, 36, 89})
	fmt.Println(l.Pop())

	fmt.Println(l.Len())

	fmt.Println(l.GetListValue())

}
