package uuid

import (
	"fmt"
	"testing"
)

func TestNewUUID(t *testing.T) {
	uintUUID, strUUID := NewUUID()

	fmt.Println(uintUUID, strUUID)
}
