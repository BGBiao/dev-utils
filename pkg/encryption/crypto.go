package encryption

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
)

func MD5(str string) string {
	hasher := md5.New()
	io.WriteString(hasher, str)
	return hex.EncodeToString(hasher.Sum(nil))

}

func HMAC(key, data string) string {
	hmac := hmac.New(md5.New, []byte(key))
	hmac.Write([]byte(data))
	return hex.EncodeToString(hmac.Sum(nil))

}

func SHA1(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum(nil))
}

func SHA256(data string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(data)))

}

func SHA512(data string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(data)))
}

func HMACSHA1(key, data string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))

}
