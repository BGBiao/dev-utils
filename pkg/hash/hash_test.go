package hash

import (
	"fmt"
	"testing"
)

func TestXXHash(t *testing.T) {

	data := "BGBiao"

	fmt.Println(XXHash(data))

}
