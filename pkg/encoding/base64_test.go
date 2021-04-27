package encoding

import (
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {

	code := Base64("BGBiao")

	fmt.Println(code)

	rawstr, err := DecodeBase64(code)

	fmt.Println(string(rawstr), err)

}
