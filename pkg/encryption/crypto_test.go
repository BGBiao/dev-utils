package encryption

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryption(t *testing.T) {
	baseStr := "BGBiao and GoOps."
	fmt.Printf("md5:%v\n", MD5(baseStr))
	fmt.Printf("hmac:%v\n", HMAC("BGBiao", baseStr))
	fmt.Printf("sha1:%v\n", SHA1(baseStr))
	fmt.Printf("sha256:%v\n", SHA256(baseStr))
	fmt.Printf("sha512:%v\n", SHA512(baseStr))
	fmt.Printf("hmacsha1:%v\n", HMACSHA1("BGBiao", baseStr))

	assert.Equal(t, MD5(baseStr), "5be3e6f61f29e151b8e23a665916eaa4i1", "md5 check ok")

}
