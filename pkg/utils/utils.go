package utils

import (
	"bytes"
	"sort"
	"strings"
	"sync"
)

var bufferPool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}

func SortedTags(tags map[string]string) string {
	if tags == nil {
		return ""
	}

	size := len(tags)

	if size == 0 {
		return ""
	}

	ret := bufferPool.Get().(*bytes.Buffer)
	ret.Reset()
	defer bufferPool.Put(ret)

	if size == 1 {
		for k, v := range tags {
			ret.WriteString(k)
			ret.WriteString("=")
			ret.WriteString(v)
		}
		return ret.String()
	}

	keys := make([]string, size)
	i := 0
	for k := range tags {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	for j, key := range keys {
		ret.WriteString(key)
		ret.WriteString("=")
		ret.WriteString(tags[key])
		if j != size-1 {
			ret.WriteString(",")
		}
	}

	return ret.String()
}

func SplitTagsString(s string) (tags map[string]string, err error) {
	err = nil
	tags = make(map[string]string)

	s = strings.Replace(s, " ", "", -1)
	if s == "" {
		return
	}

	tagSlice := strings.Split(s, ",")
	for _, tag := range tagSlice {
		tagPair := strings.SplitN(tag, "=", 2)
		if len(tagPair) == 2 {
			tags[tagPair[0]] = tagPair[1]
		} else {
			err = fmt.Errorf("bad tag %s", tag)
			return
		}
	}

	return
}

func DictedTagstring(s string) map[string]string {
	if s == "" {
		return map[string]string{}
	}
	s = strings.Replace(s, " ", "", -1)

	tag_dict := make(map[string]string)
	tags := strings.Split(s, ",")
	for _, tag := range tags {
		tag_pair := strings.SplitN(tag, "=", 2)
		if len(tag_pair) == 2 {
			tag_dict[tag_pair[0]] = tag_pair[1]
		}
	}
	return tag_dict
}
