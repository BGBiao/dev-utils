package hash

import (
	"bytes"
	"strconv"
	"sync"

	"github.com/cespare/xxhash"
)

var bufferPool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}

func XXHash(data string) string {

	ret := bufferPool.Get().(*bytes.Buffer)
	ret.Reset()
	defer bufferPool.Put(ret)

	ret.WriteString(data)

	return strconv.FormatUint(xxhash.Sum64(ret.Bytes()), 10)
}
